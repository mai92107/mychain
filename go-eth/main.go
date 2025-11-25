package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	MyToken "go-eth/go-eth/contract"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// 常數設定
	// const myTokenAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3"                        // 合約地址
	// const testNetTokenAddress = "0xc1Fd6B97B2dc53B0E93e6f3CfCC9cbE96761D8E5" // 合約地址
	// const testNetuser1Address = "0x6596D81651B271930BBc2EA5D4703159EcDe8403"                          // owner
	// const user1Address = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"                          // owner
	// const user1privateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // owner
	// const user2Address = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
	// const user2privateKey = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	// const user3Address = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
	// const user3privateKey = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
	// API_URL := "http://localhost:8545"

	// 讀取環境變數
	API_URL := os.Getenv("API_URL")
	fmt.Println(API_URL)
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	publicAddress := getPublicAddressFromPrivateKey(privateKeyHex)
	fmt.Println("Deployer address:", publicAddress.Hex())
	// user2Address := "0x160e9e408367032231b50381fe4eb10bc264eef8"
	// txHash := "0xecac96860a473d85b4313900fb880a03d3d072072763ddd556f91e096c5a3786"
	// 連線到 geth 節點
	client, err := ethclient.Dial(API_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainID := getCurrentChainID(client)
	fmt.Printf("Connected to Ethereum network with Chain ID: %s\n", chainID.String())

	// 讀取交易收據
	// readTxReceipt(client, txHash)
	
	// 部署合約到測試網路範例
	// fmt.Println("部署 MyToken 合約到測試網路...")
	// deployedAddress := deployMyTokenOnTestnet(client, privateKeyHex)
	// fmt.Printf("MyToken 合約部署完成，地址: %s\n", deployedAddress.Hex())

	// 部署合約範例
	// fmt.Println("部署 MyToken 合約...")
	// deployedAddress := deployMyToken(client, user1privateKey)
	// fmt.Printf("MyToken 合約部署完成，地址: %s\n", deployedAddress.Hex())

	// 讀取我的 Token 帳戶餘額
	// showMyTokenAccount(client, testNetTokenAddress, []string{publicAddress.Hex()})

	// 轉帳 Token 範例
	// fmt.Println("從 user1 轉帳 5 個 Token給 user2...")
	// token := getMyToken(client, testNetTokenAddress)
	// amount := big.NewInt(5 * 1e18) // 假設 token 有 18 位小數
	// tx := transferTokenByPK(token, chainID, privateKeyHex, user2Address, amount)
	// fmt.Printf("轉帳交易已送出，Tx Hash: %s\n", tx.Hash().Hex())

	// 讀取我的 Token 帳戶餘額
	// showMyTokenAccount(client, testNetTokenAddress, []string{publicAddress.Hex(), user2Address})

	// approve 範例
	// fmt.Println("user1 批准 user2 可以花費 100 個 Token...")
	// approveAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(100)) // 100 個 token
	// tx := approveTokenByPK(token, user1privateKey, user2Address, approveAmount)
	// fmt.Printf("批准交易已送出，Tx Hash: %s\n", tx.Hash().Hex())

	// 查詢 allowance 範例
	// allowance := allowanceToken(token, user1Address, user2Address)
	// fmt.Printf("user2 已被批准可以花費 user1 的 Token 數量: %s\n", formatToken(allowance, 18))

	// transferFrom 範例
	// fmt.Println("user2 從 user1 帳戶轉帳 50 個 Token 給 user3...")
	// transferFromAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(50)) // 50 個 token
	// tx = transferFromTokenByPK(token, user2privateKey, user1Address, user3Address, transferFromAmount)
	// fmt.Printf("transferFrom 交易已送出，Tx Hash: %s\n", tx.Hash().Hex())

	// 再次讀取我的 Token 帳戶餘額
	// showMyTokenAccount(client, tokenAddress, []string{user1Address, user2Address, user3Address})

	// 讀取最新區塊號
	// showLatestBlockNumber(client)

	// 啟動 HTTP 伺服器
	// showHttpServer(client, "0x5FbDB2315678afecb367f032d93F642f64180aa3")

	// showHttpServer(client, tokenAddress.Hex())
}

func showHttpServer(client *ethclient.Client, tokenAddress string) {
	r := gin.Default()

	token := getMyToken(client, tokenAddress)

	r.GET("/balance/:address", func(c *gin.Context) {
		addr := c.Param("address")

		balance := getBalance(token, addr)

		c.JSON(200, gin.H{
			"address": addr,
			"balance": balance.String(),
		})
	})

	r.POST("/account/approve", func(c *gin.Context) {

		var req struct {
			OwnerPK        string `json:"owner_pk"`
			SpenderAddress string `json:"spender_add"`
			Amount         string `json:"amt"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		reqAmt, ok := new(big.Int).SetString(req.Amount, 10)
		if !ok {
			c.JSON(400, gin.H{"error": "invalid amount"})
			return
		}

		amount := big.NewInt(0).Mul(reqAmt, big.NewInt(1e18))
		tx := approveTokenByPK(token, req.OwnerPK, req.SpenderAddress, amount)

		c.JSON(200, gin.H{
			"tx_hash": tx.Hash().Hex(),
		})
	})

	r.POST("/account/transfer", func(c *gin.Context) {
		var req struct {
			FromPK    string   `json:"from_pk"`
			ToAddress string   `json:"to_add"`
			Amount    string   `json:"amt"`
			ChainID   *big.Int `json:"chain_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		reqAmt, ok := new(big.Int).SetString(req.Amount, 10)
		if !ok {
			c.JSON(400, gin.H{"error": "invalid amount"})
			return
		}

		amount := big.NewInt(0).Mul(reqAmt, big.NewInt(1e18))
		tx := transferTokenByPK(token, req.ChainID, req.FromPK, req.ToAddress, amount)

		c.JSON(200, gin.H{
			"tx_hash": tx.Hash().Hex(),
		})
	})

	r.POST("/account/transferFrom", func(c *gin.Context) {
		var req struct {
			SpenderPK string `json:"spender_pk"`
			FromAdd   string `json:"from_add"`
			ToAddress string `json:"to_add"`
			Amount    string `json:"amt"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		reqAmt, ok := new(big.Int).SetString(req.Amount, 10)
		if !ok {
			c.JSON(400, gin.H{"error": "invalid amount"})
			return
		}

		amount := big.NewInt(0).Mul(reqAmt, big.NewInt(1e18))
		tx := transferFromTokenByPK(token, req.SpenderPK, req.FromAdd, req.ToAddress, amount)

		c.JSON(200, gin.H{
			"tx_hash": tx.Hash().Hex(),
		})
	})

	r.GET("/account/allowance/:ownerAd/:spenderAd", func(c *gin.Context) {
		ownerAd := c.Param("ownerAd")
		spenderAd := c.Param("spenderAd")

		allowance := allowanceToken(token, ownerAd, spenderAd)

		c.JSON(200, gin.H{
			"owner":     ownerAd,
			"spender":   spenderAd,
			"allowance": allowance.String(),
		})
	})

	r.Run(":8080")
}

func readTxReceipt(client *ethclient.Client, txHash string) {
	hash := common.HexToHash(txHash)
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		log.Fatalf("Failed to get tx receipt: %v", err)
	}
	printReceipt(receipt)
}

func getPublicAddressFromPrivateKey(privateKeyHex string) common.Address {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert public key type")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress
}

func getCurrentChainID(client *ethclient.Client) *big.Int {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}
	return chainID
}

func deployMyTokenOnTestnet(client *ethclient.Client, deployerPrivateKey string) common.Address {
	// 載入私鑰
	privateKey, err := crypto.HexToECDSA(deployerPrivateKey)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}
	// 建立交易簽署器
	chainID := getCurrentChainID(client)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// 設定Gas Price和Gas Limit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)

	// 部署合約
	initialSupply := new(big.Int).Mul(big.NewInt(1_000_000), big.NewInt(1e18)) // 代幣初始量
	tokenAddress, tx, _, err := MyToken.DeployMyToken(auth, client, initialSupply)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	fmt.Println("Transaction hash:", tx.Hash().Hex())

	// 🔥🔥 等待交易確認
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Failed to wait for tx: %v", err)
	}

	if receipt.Status != 1 {
		log.Fatalf("Deploy failed: receipt status = %v", receipt.Status)
	}

	fmt.Println("Contract deployed at:", tokenAddress.Hex())
	return tokenAddress
}

func deployMyTokenOnPrivateChain(client *ethclient.Client, deployerPrivateKey string) common.Address {
	ctx := context.Background()

	privateKey, _ := crypto.HexToECDSA(deployerPrivateKey)
	auth, _ := bind.NewKeyedTransactorWithChainID(
		privateKey,
		big.NewInt(31337), // 本地鏈 chain ID
	)

	// 設定Gas Price和Gas Limit
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)

	initialSupply := new(big.Int).Mul(big.NewInt(1_000_000_000), big.NewInt(1e18))

	addr, tx, tokenInstance, err := MyToken.DeployMyToken(auth, client, initialSupply) // 發行 1000000000 個 token，假設有 18 位小數
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract Address:", addr)
	fmt.Println("Deploy tx hash:", tx.Hash().Hex())

	// 等待交易確認
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deploy status:", receipt.Status)

	// 查詢餘額
	bal, _ := tokenInstance.BalanceOf(nil, auth.From)
	fmt.Println("Balance:", bal)

	return addr
}

func showMyTokenAccount(client *ethclient.Client, tokenAddress string, accounts []string) {

	// 連接到已部署的合約 -> 我發行的幣
	fmt.Println("透過 contract 部署時的 address 連接到已部署的合約（我發行的幣）...")
	token := getMyToken(client, tokenAddress)

	// 查詢 token 詳細資訊
	fmt.Println("查詢 token 詳細資訊...")
	name, symbol, decimals := getTokenDetail(token)
	fmt.Printf("Token Name: %s, Symbol: %s, Decimals: %d\n", name, symbol, decimals)

	for _, v := range accounts {
		balance := getBalance(token, v)
		fmt.Printf("Account: %s, Balance: %s\n", v, formatToken(balance, int(decimals)))
	}
}

func getTokenDetail(token *MyToken.MyToken) (string, string, uint8) {
	name, err := token.Name(nil)
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := token.Symbol(nil)
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := token.Decimals(nil)
	if err != nil {
		log.Fatal(err)
	}
	return name, symbol, decimals
}

func approveTokenByPK(token *MyToken.MyToken, ownerPrivateKey string, spenderAddress string, amount *big.Int) *types.Transaction {
	privateKey, _ := crypto.HexToECDSA(ownerPrivateKey)
	auth, _ := bind.NewKeyedTransactorWithChainID(
		privateKey,
		big.NewInt(31337), // 本地鏈 chain ID
	)
	spender := common.HexToAddress(spenderAddress)
	tx, err := token.Approve(auth, spender, amount)
	if err != nil {
		log.Fatal(err)
	}
	return tx
}

func allowanceToken(token *MyToken.MyToken, ownerAddress string, spenderAddress string) *big.Int {
	owner := common.HexToAddress(ownerAddress)
	spender := common.HexToAddress(spenderAddress)
	allowance, err := token.Allowance(nil, owner, spender)
	if err != nil {
		log.Fatal(err)
	}
	return allowance
}

func transferFromTokenByPK(token *MyToken.MyToken, spenderPrivateKey string, fromAddress string, toAddress string, amount *big.Int) *types.Transaction {
	privateKey, _ := crypto.HexToECDSA(spenderPrivateKey)
	auth, _ := bind.NewKeyedTransactorWithChainID(
		privateKey,
		big.NewInt(31337), // 本地鏈 chain ID
	)
	from := common.HexToAddress(fromAddress)
	to := common.HexToAddress(toAddress)
	tx, err := token.TransferFrom(auth, from, to, amount)
	if err != nil {
		log.Fatal(err)
	}
	return tx
}

func transferTokenByPK(token *MyToken.MyToken, chainID *big.Int, fromPrivateKey string, toAddress string, amount *big.Int) *types.Transaction {
	if chainID == nil {
		chainID = big.NewInt(31337) // 本地鏈 chain ID
	}
	privateKey, _ := crypto.HexToECDSA(fromPrivateKey)
	auth, _ := bind.NewKeyedTransactorWithChainID(
		privateKey,
		chainID,
	)
	to := common.HexToAddress(toAddress)
	tx, err := token.Transfer(auth, to, amount)
	if err != nil {
		log.Fatal(err)
	}
	return tx
}

func formatToken(balance *big.Int, decimals int) string {
	fBalance := new(big.Float).SetInt(balance)
	dec := new(big.Float).SetFloat64(1)
	for i := 0; i < decimals; i++ {
		dec.Mul(dec, big.NewFloat(10))
	}
	fBalance.Quo(fBalance, dec)
	return fBalance.Text('f', 4) // 小數點後 4 位
}

func getBalance(token *MyToken.MyToken, address string) *big.Int {
	account := common.HexToAddress(address)
	balance, err := token.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

func getMyToken(client *ethclient.Client, address string) *MyToken.MyToken {
	tokenAddress := common.HexToAddress(address)
	token, err := MyToken.NewMyToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func showLatestBlockNumber(client *ethclient.Client) {

	blockData, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// 印出區塊資訊
	printBlock(blockData)
}
func printBlock(block *types.Block) {
	fmt.Println("----------------------------------------------------")
	fmt.Println("區塊高度:", block.Number().Uint64())
	fmt.Println("區塊 Hash:", block.Hash().Hex())
	fmt.Println("Parent Hash:", block.ParentHash().Hex())
	fmt.Println("Uncle Hash:", block.UncleHash().Hex())
	fmt.Println("Tx Hash:", block.TxHash().Hex())
	fmt.Println("State Root:", block.Root().Hex())
	fmt.Println("Miner:", block.Coinbase().Hex())
	fmt.Println("時間戳:", block.Time())
	fmt.Println("Gas Limit:", block.GasLimit())
	fmt.Println("Gas Used:", block.GasUsed())
	fmt.Println("Difficulty:", block.Difficulty())
	fmt.Println("BaseFee:", block.BaseFee()) // EIP-1559 才有
	// 交易列表
	fmt.Printf("交易數量: %d\n", len(block.Transactions()))
	for _, tx := range block.Transactions() {
		printTransaction(tx)
	}

	// Uncles
	fmt.Printf("叔塊數量: %d\n", len(block.Uncles()))
	for idx, uncle := range block.Uncles() {
		fmt.Printf("叔塊 %d Hash: %s\n", idx, uncle.Hash().Hex())
	}

	// Withdrawals (如果有 Shanghai 升級)
	if block.Withdrawals() != nil {
		fmt.Printf("提款數量: %d\n", len(block.Withdrawals()))
		for _, w := range block.Withdrawals() {
			fmt.Printf("提款 %d: Validator %d, 金額 %d\n", w.Index, w.Validator, w.Amount)
		}
	}
}

func printTransaction(tx *types.Transaction) {
	fmt.Println("  ------------------- Transaction -------------------")
	fmt.Println("  Tx Hash:", tx.Hash().Hex())
	fmt.Println("  Nonce:", tx.Nonce())
	fmt.Println("  Gas Price:", tx.GasPrice())
	fmt.Println("  Gas:", tx.Gas())
	fmt.Println("  Value:", tx.Value())
	fmt.Println("  Data:", tx.Data())
	fmt.Println("  To:", func() string {
		if tx.To() == nil {
			return "Contract creation"
		}
		return tx.To().Hex()
	}())
}

func printReceipt(receipt *types.Receipt) {
	fmt.Println("----------------------------------------------------")
	fmt.Println("Receipt Tx Hash:", receipt.TxHash.Hex())
	fmt.Println("Status:", receipt.Status)
	fmt.Println("Cumulative Gas Used:", receipt.CumulativeGasUsed)
	fmt.Println("Bloom:", receipt.Bloom)
	fmt.Println("Contract Address:", func() string {
		if receipt.ContractAddress == (common.Address{}) {
			return "N/A"
		}
		return receipt.ContractAddress.Hex()
	}())
	fmt.Println("Logs:")
	for _, v := range receipt.Logs {
		fmt.Printf("  Log Index: %d, Event Creator Address: %s\n", v.Index, v.Address.Hex())

		transferSig := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
		if v.Topics[0] != transferSig {
			continue // 不是 Transfer 事件
		}
		from := common.BytesToAddress(v.Topics[1].Bytes())
		to := common.BytesToAddress(v.Topics[2].Bytes())
		value := new(big.Int).SetBytes(v.Data)
		fmt.Printf("    From: %s, To: %s, Value: %s\n", from.Hex(), to.Hex(), formatToken(value, 18))
	}
}
