import libsql_experimental as libsql
from joblib import load
import time
import smtplib
from email.mime.text import MIMEText
import os
from dotenv import load_dotenv, dotenv_values 
import resend

# loading variables from .env file
load_dotenv() 

# Define the connection parameters
url = os.getenv('TURSO_DATABASE_URL')
auth_token = os.getenv('TURSO_AUTH_TOKEN')
resend.api_key = os.getenv('RESEND_API_KEY')

# Connect to the database
conn = libsql.connect("hello.db", sync_url=url, auth_token=auth_token)
conn.sync()

# Function to send alerts
def alert(sensor_idt):
    sensor_id = sensor_idt[0]
    print(sensor_id)
    # Fetch latest sensor data for the given sensor ID
    result = conn.execute(f"SELECT do2, temperature, ph FROM sensor_data WHERE id = '{sensor_id}' AND timestamp = (SELECT MAX(timestamp) FROM sensor_data WHERE id = '{sensor_id}');").fetchone()
    
    if result:
        ph, do2, temp = result

        # Load the machine learning model
        model = load('model.joblib')

        # Make predictions
        predictions = model.predict([[ph, do2, temp]])

        # Check predictions and trigger alert if necessary
        if predictions == [-1]:
            try:
                r = resend.Emails.send({
                    "from": "waterqualiyalert@resend.dev",
                    "to": "parikshithegde2@gmail.com",
                    "subject": "Anomaly detected",
                    "html": f"<p>Anomaly detected for sensor ID {sensor_id}. <strong>Please take necessary action!</strong>!</p>"
                })
                print(f"Alert sent for sensor ID {sensor_id}")
            except Exception as e:
                print(f"Failed to send email for sensor ID {sensor_id}: {e}")
    else:
        print(f"No data found for sensor ID {sensor_id}")

# Main loop
while True:
    # Fetch distinct sensor IDs
    sensor_ids = conn.execute('SELECT DISTINCT id FROM sensor_data').fetchall()
    # Iterate over each sensor ID
    for sensor_id in sensor_ids:
        alert(sensor_id)
        conn.execute("INSERT INTO alerts (alert) VALUES ('Anomaly detected. Please take necessary action!')")
        conn.execute('commit')
    # Wait for 10 minutes before fetching new data
    time.sleep(600)  # Wait for 600 seconds (10 minutes) before fetching new data
