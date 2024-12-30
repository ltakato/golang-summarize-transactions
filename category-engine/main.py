from flask import Flask, request
import base64

app = Flask(__name__)

@app.route('/user-categories-pubsub', methods=['POST'])
def pubsub_trigger():
    """HTTP Cloud Function for Cloud Pub/Sub trigger"""
    envelope = request.get_json()
    if not envelope:
        return 'Bad Request: No Pub/Sub message received', 400

    # The Pub/Sub message data is in 'data'
    pubsub_message = envelope.get('message', {})
    data = pubsub_message.get('data', '')
    if data:
        # Decode the Pub/Sub message
        data = base64.b64decode(data).decode('utf-8')
        return f"Received message: {data}", 200
    else:
        return 'Bad Request: No message data found', 400

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)
