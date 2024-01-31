from py_algorand_sdk import account, algod, transaction

# Replace these values with your Algorand node information
ALGOD_ADDRESS = "http://localhost:4001"
ALGOD_TOKEN = "your-algod-token"

# Replace with your account information
SENDER_ADDRESS = "your-sender-address"
SENDER_PRIVATE_KEY = "your-sender-private-key"

# Replace with the recipient's address
RECIPIENT_ADDRESS = "recipient-address"

def create_algorand_client():
    algod_client = algod.AlgodClient(ALGOD_TOKEN, ALGOD_ADDRESS)
    return algod_client

def create_account():
    private_key, address = account.generate_account()
    return private_key, address

def check_balance(algod_client, address):
    account_info = algod_client.account_info(address)
    balance = account_info.get("amount")
    return balance

def send_transaction(algod_client, sender_private_key, sender_address, recipient_address, amount):
    params = algod_client.suggested_params()
    note = "Hello from Algorand!"

    unsigned_txn = transaction.PaymentTxn(sender_address, params, recipient_address, amount, None, note.encode("utf-8"))
    signed_txn = unsigned_txn.sign(sender_private_key)

    try:
        transaction_id = algod_client.send_transaction(signed_txn)
        print(f"Transaction ID: {transaction_id}")
    except Exception as e:
        print(f"Error sending transaction: {e}")

def main():
    algod_client = create_algorand_client()

    # Create an account (if you don't have one)
    private_key, address = create_account()
    print(f"New Account Created: {address}")

    # Check balance
    balance = check_balance(algod_client, address)
    print(f"Account Balance: {balance} microAlgos")

    # Send a transaction (replace RECIPIENT_ADDRESS and amount)
    send_transaction(algod_client, private_key, address, RECIPIENT_ADDRESS, 1000000)

if __name__ == "__main__":
    main()
