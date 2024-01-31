from substrateinterface import SubstrateInterface, Keypair
from substrateinterface.exceptions import SubstrateRequestException

# Replace with the actual WebSocket RPC endpoint of your Substrate node
substrate_url = "ws://localhost:9944"

# Replace with the actual Alice's account seed or address
alice_seed = "Alice's Seed or Address"

def get_balance(substrate, account_address):
    try:
        balance = substrate.get_balance(account_address)
        return balance
    except SubstrateRequestException as e:
        print(f"Error: {e}")
        return None

def main():
    # Connect to the Substrate node
    substrate = SubstrateInterface(url=substrate_url)

    # Create Keypair from Alice's seed
    alice_keypair = Keypair.create_from_mnemonic(alice_seed)

    # Get Alice's account address
    alice_address = alice_keypair.ss58_address

    # Get and print Alice's account balance
    balance = get_balance(substrate, alice_address)
    if balance is not None:
        print(f"Alice's Balance: {balance} Unit")

if __name__ == "__main__":
    main()
