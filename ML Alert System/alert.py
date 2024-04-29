import libsql_experimental as libsql
from joblib import load
import requests
import time

# Define the connection parameters
url = "libsql://capstone-sibearian.turso.io"
auth_token = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MTMxOTkyMTksImlkIjoiNzNjMzJjNmYtMzhlNC00MTNlLTliZTgtMDM1M2VkNjYzZjRlIn0.oMGJmQe56pBTMt5auJja63N0QEr3i84qhbNJtvkxwdPwlz0l1RmB5C07_HVBVF7USdFvFMvweNubDOrS5rzzAg"

# Connect to the database
conn = libsql.connect("hello.db", sync_url=url, auth_token=auth_token)
conn.sync()

# Function to send alerts
def alert():
    # Define the URL and headers
    url = 'https://graph.facebook.com/295149117018021/messages'
    headers = {
        'Authorization': 'Bearer EAANLrM0xTmoBOwNGDf0uWAA1n0ZArdZAGMkHatZABZAQoZANMGwL9DNCpHJtMZBxduHbRatHZAG98JpRvGSNyu1DdoqIE3eZAdAgYJBwGYUy5n3JXhUNxsmuq9HxywkPJHNWZBlJha7ICCpAjEAKlx3BLZBnJPPE92KwZAU0jOozVHivSVI5TY7OnKBEAkEWychRJnnlWG7lqGU3ZBl4vKPn7KoD4zZAqDQ7eCKIiwHMZD',
        'Content-Type': 'application/json'
    }

    # Define the request body
    payload = {
        "messaging_product": "whatsapp",
        "recipient_type": "individual",
        "to": "919008077372",
        "type": "text",
        "text": {"body": "Testing Whatsapp custom message"}
    }

    # Send the POST request
    response = requests.post(url, headers=headers, json=payload)

    # Check if the message was sent successfully
    data = response.json()
    if 'messages' in data and len(data['messages']) > 0:
        print("Message sent successfully! Message ID:", data['messages'][0]['id'])
    else:
        print("Failed to send message.")
    
# Load the machine learning model
model = load('model.joblib')

# Main loop
while True:
    # Fetch latest sensor data
    result = conn.execute("SELECT do2, temperature, ph FROM sensor_data WHERE timestamp = (SELECT MAX(timestamp) FROM sensor_data);").fetchone()
    if result:
        ph = result[2]
        do2 = result[0]
        temp = result[1]

        # Make predictions
        predictions = model.predict([[ph, do2, temp]])

        # Check predictions and trigger alert if necessary
        if predictions == [-1]:
            #alert()
            conn.execute("INSERT INTO alerts (alert) VALUES ('tests')")
            conn.execute('commit')
    else:
        print("No sensor data found.")
    
    # Wait for 10 minutes before fetching new data
    time.sleep(600)  # Wait for 600 seconds (10 minutes) before fetching new data
