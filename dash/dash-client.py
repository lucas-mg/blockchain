from dash import Dash, PrivateKey

# Replace with your actual private key
PRIVATE_KEY = "your_private_key_here"

def create_dash_client():
    return Dash()

def create_wallet():
    private_key = PrivateKey(PRIVATE_KEY)
    address = private_key.address
    return private_key, address

def check_balance(dash_client, address):
    balance = dash_client.get_balance(address)
    return balance

def send_transaction(dash_client, sender_private_key, recipient_address, amount):
    transaction = dash_client.send_to_address(recipient_address, amount, sender_private_key=sender_private_key)
    return transaction

def main():
    dash_client = create_dash_client()

    # Create a wallet (if you don't have one)
    private_key, address = create_wallet()
    print(f"New Wallet Created: {address}")

    # Check balance
    balance = check_balance(dash_client, address)
    print(f"Wallet Balance: {balance} DASH")

    # Send a transaction (replace RECIPIENT_ADDRESS and amount)
    recipient_address = "recipient-address"
    amount = 0.1
    transaction = send_transaction(dash_client, private_key, recipient_address, amount)
    print(f"Transaction ID: {transaction['txid']}")

if __name__ == "__main__":
    main()
