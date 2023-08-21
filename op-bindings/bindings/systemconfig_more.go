// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"espresso\",\"offset\":8,\"slot\":\"104\",\"type\":\"t_bool\"},{\"astId\":1010,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1012_storage\"},{\"astId\":1011,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"startBlock\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b506004361061025c5760003560e01c8063935f029e11610145578063dac6e63a116100bd578063f45e65d81161008c578063f8c68de011610071578063f8c68de0146105b3578063fd32aa0f146105bb578063ffa1ad74146105c357600080fd5b8063f45e65d814610596578063f68016b71461059f57600080fd5b8063dac6e63a1461055f578063e81b2c6d14610567578063f06d7a3b14610570578063f2fde38b1461058357600080fd5b8063bc49ce5f11610114578063c71973f6116100f9578063c71973f614610405578063c9b26f6114610418578063cc731b021461042b57600080fd5b8063bc49ce5f146103f5578063c4e8ddfa146103fd57600080fd5b8063935f029e146103bf5780639b7d7f0a146103d2578063a7119869146103da578063b40a817c146103e257600080fd5b80634add321d116101d85780635d73369c116101a757806361fba0ca1161018c57806361fba0ca14610370578063715018a6146103995780638da5cb5b146103a157600080fd5b80635d73369c1461036057806361d157681461036857600080fd5b80634add321d146102fb5780634d9f15591461031c5780634f16540b1461032457806354fd4d501461034b57600080fd5b806318d139181161022f5780631fd19ee1116102145780631fd19ee1146102d757806336eeb0e6146102df57806348cd4cb1146102f257600080fd5b806318d13918146102ba57806319f5cea8146102cf57600080fd5b806306c9265714610261578063078f29cf1461027c5780630a49cb03146102a95780630c18c162146102b1575b600080fd5b6102696105cb565b6040519081526020015b60405180910390f35b6102846105f9565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610273565b610284610632565b61026960655481565b6102cd6102c836600461199a565b610662565b005b610269610676565b6102846106a1565b6102cd6102ed366004611b69565b6106cb565b610269606a5481565b610303610a58565b60405167ffffffffffffffff9091168152602001610273565b610284610a7e565b6102697f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b610353610aae565b6040516102739190611d2c565b610269610b51565b610269610b7c565b6068546103899068010000000000000000900460ff1681565b6040519015158152602001610273565b6102cd610ba7565b60335473ffffffffffffffffffffffffffffffffffffffff16610284565b6102cd6103cd366004611d3f565b610bbb565b610284610bd1565b610284610c01565b6102cd6103f0366004611d61565b610c31565b610269610c42565b610284610c6d565b6102cd610413366004611d7c565b610c9d565b6102cd610426366004611d98565b610cae565b6104ef6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b6040516102739190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b610284610cbf565b61026960675481565b6102cd61057e366004611db1565b610cef565b6102cd61059136600461199a565b610d00565b61026960665481565b6068546103039067ffffffffffffffff1681565b610269610db4565b610269610ddf565b610269600081565b6105f660017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611dfb565b81565b600061062d61062960017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611dfb565b5490565b905090565b600061062d61062960017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611dfb565b61066a610e0a565b61067381610e8b565b50565b6105f660017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611dfb565b600061062d7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c085490565b600054600290610100900460ff161580156106ed575060005460ff8083169116105b61077e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff8316176101001790556107b7610f47565b6107c08c610d00565b6107c989610fe6565b6107d38b8b61100e565b6107dc8861109f565b6107e586610e8b565b6107ee8761117d565b6108208361081d60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611dfb565b55565b81516108519061081d60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611dfb565b60208201516108859061081d60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611dfb565b60408201516108b99061081d60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611dfb565b60608201516108ed9061081d60017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611dfb565b60808201516109219061081d60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611dfb565b60a08201516109559061081d60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611dfb565b61095e84611207565b610967856112a9565b61096f610a58565b67ffffffffffffffff168867ffffffffffffffff1610156109ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610775565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050505050565b60695460009061062d9063ffffffff6a0100000000000000000000820481169116611e12565b600061062d61062960017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611dfb565b6060610ad97f000000000000000000000000000000000000000000000000000000000000000061171d565b610b027f000000000000000000000000000000000000000000000000000000000000000061171d565b610b2b7f000000000000000000000000000000000000000000000000000000000000000061171d565b604051602001610b3d93929190611e3e565b604051602081830303815290604052905090565b6105f660017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611dfb565b6105f660017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611dfb565b610baf610e0a565b610bb9600061185a565b565b610bc3610e0a565b610bcd828261100e565b5050565b600061062d61062960017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611dfb565b600061062d61062960017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611dfb565b610c39610e0a565b6106738161109f565b6105f660017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611dfb565b600061062d61062960017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611dfb565b610ca5610e0a565b610673816112a9565b610cb6610e0a565b61067381610fe6565b600061062d61062960017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611dfb565b610cf7610e0a565b6106738161117d565b610d08610e0a565b73ffffffffffffffffffffffffffffffffffffffff8116610dab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610775565b6106738161185a565b6105f660017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611dfb565b6105f660017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611dfb565b60335473ffffffffffffffffffffffffffffffffffffffff163314610bb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610775565b610eb3817f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0855565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be83604051610f3b9190611d2c565b60405180910390a35050565b600054610100900460ff16610fde576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610775565b610bb96118d1565b6067819055604080516020808201849052825180830390910181529082019091526000610f0a565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516110929190611d2c565b60405180910390a3505050565b6110a7610a58565b67ffffffffffffffff168167ffffffffffffffff161015611124576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610775565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002610f0a565b6068805482151568010000000000000000027fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff9091161790556040516000906111d0908390602001901515815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905290506004610f0a565b606a5415611297576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603860248201527f53797374656d436f6e6669673a2063616e6e6f74206f7665727269646520616e60448201527f20616c72656164792073657420737461727420626c6f636b00000000000000006064820152608401610775565b80156112a257606a55565b43606a5550565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff161115611359576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d6178206261736500000000000000000000006064820152608401610775565b6001816040015160ff16116113f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e203100000000000000000000000000000000006064820152608401610775565b6068546080820151825167ffffffffffffffff909216916114119190611eb4565b63ffffffff16111561147f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610775565b6000816020015160ff1611611516576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f74206265203000000000000000000000000000000000006064820152608401610775565b8051602082015163ffffffff82169160ff90911690611536908290611f02565b6115409190611f25565b63ffffffff16146115d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d69740000000000000000006064820152608401610775565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b60608160000361176057505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561178a578061177481611f51565b91506117839050600a83611f89565b9150611764565b60008167ffffffffffffffff8111156117a5576117a56119e4565b6040519080825280601f01601f1916602001820160405280156117cf576020820181803683370190505b5090505b8415611852576117e4600183611dfb565b91506117f1600a86611f9d565b6117fc906030611fb1565b60f81b81838151811061181157611811611fc9565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061184b600a86611f89565b94506117d3565b949350505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16611968576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610775565b610bb93361185a565b803573ffffffffffffffffffffffffffffffffffffffff8116811461199557600080fd5b919050565b6000602082840312156119ac57600080fd5b6119b582611971565b9392505050565b803567ffffffffffffffff8116811461199557600080fd5b8035801515811461199557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160c0810167ffffffffffffffff81118282101715611a5d577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b803563ffffffff8116811461199557600080fd5b803560ff8116811461199557600080fd5b600060c08284031215611a9a57600080fd5b60405160c0810181811067ffffffffffffffff82111715611ae4577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052905080611af383611a63565b8152611b0160208401611a77565b6020820152611b1260408401611a77565b6040820152611b2360608401611a63565b6060820152611b3460808401611a63565b608082015260a08301356fffffffffffffffffffffffffffffffff81168114611b5c57600080fd5b60a0919091015292915050565b60008060008060008060008060008060008b8d036102a0811215611b8c57600080fd5b611b958d611971565b9b5060208d01359a5060408d0135995060608d01359850611bb860808e016119bc565b9750611bc660a08e016119d4565b9650611bd460c08e01611971565b9550611be38e60e08f01611a88565b94506101a08d01359350611bfa6101c08e01611971565b925060c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe2082011215611c2c57600080fd5b50611c35611a13565b611c426101e08e01611971565b8152611c516102008e01611971565b6020820152611c636102208e01611971565b6040820152611c756102408e01611971565b6060820152611c876102608e01611971565b6080820152611c996102808e01611971565b60a0820152809150509295989b509295989b9093969950565b60005b83811015611ccd578181015183820152602001611cb5565b83811115611cdc576000848401525b50505050565b60008151808452611cfa816020860160208601611cb2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006119b56020830184611ce2565b60008060408385031215611d5257600080fd5b50508035926020909101359150565b600060208284031215611d7357600080fd5b6119b5826119bc565b600060c08284031215611d8e57600080fd5b6119b58383611a88565b600060208284031215611daa57600080fd5b5035919050565b600060208284031215611dc357600080fd5b6119b5826119d4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611e0d57611e0d611dcc565b500390565b600067ffffffffffffffff808316818516808303821115611e3557611e35611dcc565b01949350505050565b60008451611e50818460208901611cb2565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611e8c816001850160208a01611cb2565b60019201918201528351611ea7816002840160208801611cb2565b0160020195945050505050565b600063ffffffff808316818516808303821115611e3557611e35611dcc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600063ffffffff80841680611f1957611f19611ed3565b92169190910492915050565b600063ffffffff80831681851681830481118215151615611f4857611f48611dcc565b02949350505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611f8257611f82611dcc565b5060010190565b600082611f9857611f98611ed3565b500490565b600082611fac57611fac611ed3565b500690565b60008219821115611fc457611fc4611dcc565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
