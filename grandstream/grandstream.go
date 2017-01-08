package grandstream

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type keyCode struct {
	Account int
	Name    int
	UserId  int
}

var keyCodes = map[int]keyCode{
// phone-builtin speed-dial buttons
	0:  keyCode{301, 302, 303},
	1:  keyCode{304, 305, 306},
	2:  keyCode{307, 308, 309},
	3:  keyCode{310, 311, 312},
	4:  keyCode{313, 314, 315},
	5:  keyCode{316, 317, 318},
	6:  keyCode{319, 320, 321},
	7:  keyCode{354, 355, 356},
	8:  keyCode{358, 359, 360},
	9:  keyCode{362, 363, 364},
	10: keyCode{366, 367, 368},
	11: keyCode{370, 371, 372},
	12: keyCode{374, 375, 376},
	13: keyCode{378, 379, 380},
	14: keyCode{382, 383, 384},
	15: keyCode{386, 387, 388},
	16: keyCode{390, 391, 392},
	17: keyCode{394, 395, 396},

// 1st extension panel
	1024 + 0:  keyCode{6201, 6401, 6601},
	1024 + 1:  keyCode{6202, 6402, 6602},
	1024 + 2:  keyCode{6203, 6403, 6603},
	1024 + 3:  keyCode{6204, 6404, 6604},
	1024 + 4:  keyCode{6205, 6405, 6605},
	1024 + 5:  keyCode{6206, 6406, 6606},
	1024 + 6:  keyCode{6207, 6407, 6607},
	1024 + 7:  keyCode{6208, 6408, 6608},
	1024 + 8:  keyCode{6209, 6409, 6609},
	1024 + 9:  keyCode{6210, 6410, 6610},
	1024 + 10: keyCode{6211, 6411, 6611},
	1024 + 11: keyCode{6212, 6412, 6612},
	1024 + 12: keyCode{6213, 6413, 6613},
	1024 + 13: keyCode{6214, 6414, 6614},
	1024 + 14: keyCode{6215, 6415, 6615},
	1024 + 15: keyCode{6216, 6416, 6616},
	1024 + 16: keyCode{6217, 6417, 6617},
	1024 + 17: keyCode{6218, 6418, 6618},
	1024 + 18: keyCode{6219, 6419, 6619},
	1024 + 19: keyCode{6220, 6420, 6620},
	1024 + 20: keyCode{6221, 6421, 6621},
	1024 + 21: keyCode{6222, 6422, 6622},
	1024 + 22: keyCode{6223, 6423, 6623},
	1024 + 23: keyCode{6224, 6424, 6624},
	1024 + 24: keyCode{6225, 6425, 6625},
	1024 + 25: keyCode{6226, 6426, 6626},
	1024 + 26: keyCode{6227, 6427, 6627},
	1024 + 27: keyCode{6228, 6428, 6628},
	1024 + 28: keyCode{6229, 6429, 6629},
	1024 + 29: keyCode{6230, 6430, 6630},
	1024 + 30: keyCode{6231, 6431, 6631},
	1024 + 31: keyCode{6232, 6432, 6632},
	1024 + 32: keyCode{6233, 6433, 6633},
	1024 + 33: keyCode{6234, 6434, 6634},
	1024 + 34: keyCode{6235, 6435, 6635},
	1024 + 35: keyCode{6236, 6436, 6636},
	1024 + 36: keyCode{6237, 6437, 6637},
	1024 + 37: keyCode{6238, 6438, 6638},
	1024 + 38: keyCode{6239, 6439, 6639},
	1024 + 39: keyCode{6240, 6440, 6640},
	1024 + 40: keyCode{6241, 6441, 6641},
	1024 + 41: keyCode{6242, 6442, 6642},
	1024 + 42: keyCode{6243, 6443, 6643},
	1024 + 43: keyCode{6244, 6444, 6644},
	1024 + 44: keyCode{6245, 6445, 6645},
	1024 + 45: keyCode{6246, 6446, 6646},
	1024 + 46: keyCode{6247, 6447, 6647},
	1024 + 47: keyCode{6248, 6448, 6648},
	1024 + 48: keyCode{6249, 6449, 6649},
	1024 + 49: keyCode{6250, 6450, 6650},
	1024 + 50: keyCode{6251, 6451, 6651},
	1024 + 51: keyCode{6252, 6452, 6652},
	1024 + 52: keyCode{6253, 6453, 6653},
	1024 + 53: keyCode{6254, 6454, 6654},
	1024 + 54: keyCode{6255, 6455, 6655},
	1024 + 55: keyCode{6256, 6456, 6656},

// 2nd extension panel
	2048 + 0:  keyCode{6257, 6457, 6657},
	2048 + 1:  keyCode{6258, 6458, 6658},
	2048 + 2:  keyCode{6259, 6459, 6659},
	2048 + 3:  keyCode{6260, 6460, 6660},
	2048 + 4:  keyCode{6261, 6461, 6661},
	2048 + 5:  keyCode{6262, 6462, 6662},
	2048 + 6:  keyCode{6263, 6463, 6663},
	2048 + 7:  keyCode{6264, 6464, 6664},
	2048 + 8:  keyCode{6265, 6465, 6665},
	2048 + 9:  keyCode{6266, 6466, 6666},
	2048 + 10: keyCode{6267, 6467, 6667},
	2048 + 11: keyCode{6268, 6468, 6668},
	2048 + 12: keyCode{6269, 6469, 6669},
	2048 + 13: keyCode{6270, 6470, 6670},
	2048 + 14: keyCode{6271, 6471, 6671},
	2048 + 15: keyCode{6272, 6472, 6672},
	2048 + 16: keyCode{6273, 6473, 6673},
	2048 + 17: keyCode{6274, 6474, 6674},
	2048 + 18: keyCode{6275, 6475, 6675},
	2048 + 19: keyCode{6276, 6476, 6676},
	2048 + 20: keyCode{6277, 6477, 6677},
	2048 + 21: keyCode{6278, 6478, 6678},
	2048 + 22: keyCode{6279, 6479, 6679},
	2048 + 23: keyCode{6280, 6480, 6680},
	2048 + 24: keyCode{6281, 6481, 6681},
	2048 + 25: keyCode{6282, 6482, 6682},
	2048 + 26: keyCode{6283, 6483, 6683},
	2048 + 27: keyCode{6284, 6484, 6684},
	2048 + 28: keyCode{6285, 6485, 6685},
	2048 + 29: keyCode{6286, 6486, 6686},
	2048 + 30: keyCode{6287, 6487, 6687},
	2048 + 31: keyCode{6288, 6488, 6688},
	2048 + 32: keyCode{6289, 6489, 6689},
	2048 + 33: keyCode{6290, 6490, 6690},
	2048 + 34: keyCode{6291, 6491, 6691},
	2048 + 35: keyCode{6292, 6492, 6692},
	2048 + 36: keyCode{6293, 6493, 6693},
	2048 + 37: keyCode{6294, 6494, 6694},
	2048 + 38: keyCode{6295, 6495, 6695},
	2048 + 39: keyCode{6296, 6496, 6696},
	2048 + 40: keyCode{6297, 6497, 6697},
	2048 + 41: keyCode{6298, 6498, 6698},
	2048 + 42: keyCode{6299, 6499, 6699},
	2048 + 43: keyCode{6300, 6500, 6700},
	2048 + 44: keyCode{6301, 6501, 6701},
	2048 + 45: keyCode{6302, 6502, 6702},
	2048 + 46: keyCode{6303, 6503, 6703},
	2048 + 47: keyCode{6304, 6504, 6704},
	2048 + 48: keyCode{6305, 6505, 6705},
	2048 + 49: keyCode{6306, 6506, 6706},
	2048 + 50: keyCode{6307, 6507, 6707},
	2048 + 51: keyCode{6308, 6508, 6708},
	2048 + 52: keyCode{6309, 6509, 6709},
	2048 + 53: keyCode{6310, 6510, 6710},
	2048 + 54: keyCode{6311, 6511, 6711},
	2048 + 55: keyCode{6312, 6512, 6712},
}

type fullEntryName struct {
	CategoryName string
	SubcategoryName string
	EntryName string
}

func ParseFile(in *os.File) (map[string]map[string]map[string]string, error) {
	entryNameMap := make(map[int]fullEntryName)
	for keyId, keyCode := range keyCodes {
		extensionModuleId := int(keyId / 1024)
		var categoryName string
		switch extensionModuleId {
		case 0:
			categoryName = "BuiltinSpeedDialButtons"
		default:
			categoryName = fmt.Sprintf("ExtensionModule%vSpeedDialButtons", extensionModuleId)
		}

		subcategoryName := fmt.Sprintf("SpeedDialButton%v", keyId % 1024)
		
		entryNameMap[keyCode.Account] = fullEntryName{categoryName, subcategoryName, "Account"}
		entryNameMap[keyCode.Name   ] = fullEntryName{categoryName, subcategoryName, "Name"   }
		entryNameMap[keyCode.UserId ] = fullEntryName{categoryName, subcategoryName, "UserId" }
	}

	result := make(map[string]map[string]map[string]string)

	reader := bufio.NewReader(in)

	for {
		lineBytes, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if isPrefix {
			return nil, fmt.Errorf("line is too long")
		}
		line := string(lineBytes)
		words := strings.Split(line, "=")
		pCodeStr := words[0]
		arg := words[1]

		if pCodeStr[0:1] != "P" {
			return nil, fmt.Errorf("Invalid P-Code: \"%v\"", pCodeStr)
		}

		pCode, err := strconv.Atoi(pCodeStr[1:])
		if err != nil {
			return nil, err
		}

		fullEntryN, ok := entryNameMap[pCode]
		if !ok {
			return nil, fmt.Errorf("Unknown P-Code: %v. Known P-Codes: %v", pCode, entryNameMap)
		}

		if result[fullEntryN.CategoryName] == nil {
			result[fullEntryN.CategoryName] = make(map[string]map[string]string)
		}
		if result[fullEntryN.CategoryName][fullEntryN.SubcategoryName] == nil {
			result[fullEntryN.CategoryName][fullEntryN.SubcategoryName] = make(map[string]string)
		}

		result[fullEntryN.CategoryName][fullEntryN.SubcategoryName][fullEntryN.EntryName] = arg
	}

	return result, nil
}
func WriteToFile(in *os.File, configuration map[string]map[string]map[string]string) error {
	return nil
}
