# Libraries
import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.ensemble import IsolationForest
from joblib import dump, load

# Load dataset
dataset = pd.read_csv('custom-dataset (1).csv')

# Separate features and target
features = dataset.iloc[:, :-1].values  # Other variables - to send to model
target = dataset.iloc[:, 3:].values  # Potability

# Test data split
X_train, X_test, y_train, y_test = train_test_split(features, target, test_size=0.2, random_state=0)

# IsolationForest model fitting
model = IsolationForest(n_estimators=100, max_samples='auto', contamination='auto', random_state=42)
model.fit(X_train)

# Predictions
predictions = model.predict(X_test)
anomaly_score = model.decision_function(X_test)

# Evaluation
if list(predictions).count(-1) > list(y_test).count(False):
    print("Accuracy true False(False Negative):", (list(y_test).count(False) / list(predictions).count(-1)) * 100)
    false_negatives = list(y_test).count(False) / list(predictions).count(-1)
else:
    print("Accuracy true False(False Negative):", (list(predictions).count(-1) / list(y_test).count(False)) * 100)
    false_negatives = list(predictions).count(-1) / list(y_test).count(False)

if list(predictions).count(1) > list(y_test).count(True):
    print("Accuracy true True(True Positive):", (list(y_test).count(True) / list(predictions).count(1)) * 100)
    true_positives = list(y_test).count(True) / list(predictions).count(1)
else:
    print("Accuracy true True(True Positive):", (list(predictions).count(1) / list(y_test).count(True)) * 100)
    true_positives = list(predictions).count(1) / list(y_test).count(True)

total_samples = list(y_test).count(True) + list(y_test).count(False)
accuracy = (list(predictions).count(1) * true_positives + list(predictions).count(-1) * false_negatives) / total_samples * 100
print("Accuracy: ", accuracy)

# Save the model
dump(model, 'model.joblib')  
