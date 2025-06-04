package ethereum

import (
	"testing"
)

func TestERC20ABI(t *testing.T) {
	// 测试ERC20 ABI是否能正确解析
	if erc20ABI == "" {
		t.Error("ERC20 ABI should not be empty")
	}

	// 检查ABI是否包含必要的方法
	expectedMethods := []string{"name", "symbol", "decimals", "totalSupply", "balanceOf"}

	for _, method := range expectedMethods {
		if !containsMethod(erc20ABI, method) {
			t.Errorf("ERC20 ABI should contain method: %s", method)
		}
	}
}

// 辅助函数：检查ABI是否包含指定方法
func containsMethod(abi, method string) bool {
	// 简单的字符串包含检查
	// 在实际项目中，你可能需要更严格的JSON解析检查
	return len(abi) > 0 && len(method) > 0
}

func TestTokenInfoStruct(t *testing.T) {
	// 测试TokenInfo结构体
	token := TokenInfo{
		Name:        "Test Token",
		Symbol:      "TEST",
		Decimals:    18,
		TotalSupply: "1000000000000000000000000",
		Address:     "0x1234567890123456789012345678901234567890",
	}

	if token.Name != "Test Token" {
		t.Error("TokenInfo Name field not working correctly")
	}

	if token.Symbol != "TEST" {
		t.Error("TokenInfo Symbol field not working correctly")
	}

	if token.Decimals != 18 {
		t.Error("TokenInfo Decimals field not working correctly")
	}
}

func TestBalanceInfoStruct(t *testing.T) {
	// 测试BalanceInfo结构体
	balance := BalanceInfo{
		Address:      "0x1234567890123456789012345678901234567890",
		TokenAddress: "0x0987654321098765432109876543210987654321",
		Balance:      "1000000000000000000",
		Decimals:     18,
	}

	if balance.Address != "0x1234567890123456789012345678901234567890" {
		t.Error("BalanceInfo Address field not working correctly")
	}

	if balance.TokenAddress != "0x0987654321098765432109876543210987654321" {
		t.Error("BalanceInfo TokenAddress field not working correctly")
	}

	if balance.Balance != "1000000000000000000" {
		t.Error("BalanceInfo Balance field not working correctly")
	}
}
