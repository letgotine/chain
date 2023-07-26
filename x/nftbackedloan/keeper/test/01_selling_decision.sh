#!/bin/sh

# block speed
sleep=5

init_user1_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom uguu -o json | jq .amount | tr -d '"')
init_user2_balance=$(ununifid q bank balances ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla --denom uguu -o json | jq .amount | tr -d '"')

# mint nft
echo "------------mint nft------------"
ununifid tx nftfactory mint-nft \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# list nft
echo "------------list nft------------"
ununifid tx nftbackedloan list \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# bid nft from user2
echo "------------bid nft from user2------------"
ununifid tx nftbackedloan place-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
200uguu 50uguu 0.1 7200 --automatic-payment=false --from user2 --keyring-backend test --chain-id test --yes

sleep $sleep
# selling decision
echo "------------selling decision------------"
ununifid tx nftbackedloan selling-decision \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01

# pay remainder
echo "------------pay remainder------------"
ununifid tx nftbackedloan pay-remainder \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

# Check nft transfer & balance
user2_balance=$(ununifid q bank balances ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla --denom uguu -o json | jq .amount | tr -d '"')

# -200
expected_user2_balance=$(($init_user2_balance - 200))

if [ "$user2_balance" = "$expected_user2_balance" ]; then
  echo "pass: bidder balance is correct: $user2_balance"
else
  echo "error: bidder balance is incorrect: $user2_balance"
fi

sleep $sleep
echo "wait.......... NftListingFullPaymentPeriod: 30"
sleep 30

echo "============check nft status successful_bid ============"
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01
echo "wait.......... NftListingNftDeliveryPeriod: 30s"
sleep 30

user1_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom uguu -o json | jq .amount | tr -d '"')
# + 200 - 10 (5% fee)
expected_user1_balance=$(($init_user1_balance + 190))

if [ "$user1_balance" = "$expected_user1_balance" ]; then
  echo "pass: owner balance is correct: $user1_balance"
else
  echo "error: owner balance is incorrect:"
  echo "expected: $expected_user1_balance actual: $user1_balance"
fi

owner=$(ununifid q nft owner ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 -o json | jq .owner | tr -d '"')
# user2
expected_owner="ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla"

if [ "$owner" = "$expected_owner" ]; then
  echo "pass: Owner is changed to $owner"
else
  echo "error: Owner is not changed from $owner"
fi