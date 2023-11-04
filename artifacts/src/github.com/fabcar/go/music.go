package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

// SmartContract Define the Smart Contract structure
type SmartContract struct {
}

// Artiste :  Define the Artiste structure, with 4 properties.  Structure tags are used by encoding/json library
type Album struct {
	AlbumName   string `json:"albumname"`
	Ext  string `json:"ext"`
	Title string `json:"title"`
	Owner  string `json:"owner"`
}

type artistePrivateDetails struct {
	Owner string `json:"owner"`
	Price string `json:"price"`
}

// Init ;  Method for initializing smart contract
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

var logger = flogging.MustGetLogger("fabArtiste_cc")

// Invoke :  Method for INVOKING smart contract
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	logger.Infof("Function name is:  %d", function)
	logger.Infof("Args length is : %d", len(args))

	switch function {
	case "queryAlbum":
		return s.queryAlbum(APIstub, args)
	case "initLedger":
		return s.initLedger(APIstub)
	case "createAlbum":
		return s.createAlbum(APIstub, args)
	case "queryAllAlbums":
		return s.queryAllAlbums(APIstub)
	case "changeAlbumOwner":
		return s.changeAlbumOwner(APIstub, args)
	case "getHistoryForAsset":
		return s.getHistoryForAsset(APIstub, args)
	case "queryAlbumsByOwner":
		return s.queryAlbumsByOwner(APIstub, args)
	case "restictedMethod":
		return s.restictedMethod(APIstub, args)
	case "test":
		return s.test(APIstub, args)
	case "createPrivateAlbum":
		return s.createPrivateAlbum(APIstub, args)
	case "readPrivateAlbum":
		return s.readPrivateAlbum(APIstub, args)
	case "updatePrivateAlbum":
		return s.updatePrivateAlbum(APIstub, args)
	case "readAlbumPrivateDetails":
		return s.readAlbumPrivateDetails(APIstub, args)
	case "createPrivateAlbumImplicitForOrg1":
		return s.createPrivateAlbumImplicitForOrg1(APIstub, args)
	case "createPrivateAlbumImplicitForOrg2":
		return s.createPrivateAlbumImplicitForOrg2(APIstub, args)
	case "queryPrivateDataHash":
		return s.queryPrivateDataHash(APIstub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}

	// return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryAlbum(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ArtisteAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(ArtisteAsBytes)
}

func (s *SmartContract) readPrivateAlbum(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	// collectionArtistes, collectionArtistePrivateDetails, _implicit_org_Org1MSP, _implicit_org_Org2MSP
	ArtisteAsBytes, err := APIstub.GetPrivateData(args[0], args[1])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get private details for " + args[1] + ": " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	} else if ArtisteAsBytes == nil {
		jsonResp := "{\"Error\":\"Artiste private details does not exist: " + args[1] + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(ArtisteAsBytes)
}

func (s *SmartContract) readPrivateArtisteIMpleciteForOrg1(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ArtisteAsBytes, _ := APIstub.GetPrivateData("_implicit_org_Org1MSP", args[0])
	return shim.Success(ArtisteAsBytes)
}

func (s *SmartContract) readAlbumPrivateDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ArtisteAsBytes, err := APIstub.GetPrivateData("collectionArtistePrivateDetails", args[0])

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get private details for " + args[0] + ": " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	} else if ArtisteAsBytes == nil {
		jsonResp := "{\"Error\":\"Marble private details does not exist: " + args[0] + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(ArtisteAsBytes)
}

func (s *SmartContract) test(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ArtisteAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(ArtisteAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	artistes := []Album{
		Album{AlbumName: "OverdemAll", Ext: ".mp3", Title: "blue", Owner: "Davido"},
		Album{AlbumName: "Damini", Ext: ".mp3", Title: "red", Owner: "Wizkid"},
		Album{AlbumName: "OverdemAll", Ext: ".mp3", Title: "green", Owner: "Davido"},
		Album{AlbumName: "Damini", Ext: ".mp3", Title: "yellow", Owner: "Wizkid"},
		Album{AlbumName: "OverdemAll", Ext: ".mp3", Title: "black", Owner: "Davido"},
		Album{AlbumName: "Damini", Ext: ".mp3", Title: "purple", Owner: "Wizkid"},
		Album{AlbumName: "OverdemAll", Ext: ".mp3", Title: "white", Owner: "Davido"},
		Album{AlbumName: "Damini", Ext: ".mp3", Title: "violet", Owner: "Wizkid"},
		Album{AlbumName: "OverdemAll", Ext: ".mp3", Title: "indigo", Owner: "Davido"},
		Album{AlbumName: "Damini", Ext: ".mp3", Title: "brown", Owner: "Wizkid"},
	}

	i := 0
	for i < len(artistes) {
		ArtisteAsBytes, _ := json.Marshal(artistes[i])
		APIstub.PutState("Artiste"+strconv.Itoa(i), ArtisteAsBytes)
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createPrivateAlbum(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	type ArtisteTransientInput struct {
		Album  string `json:"album"` //the fieldtags are needed to keep case from bouncing around
		Ext string `json:"ext"`
		Title string `json:"title"`
		Owner string `json:"owner"`
		Price string `json:"price"`
		Key   string `json:"key"`
	}
	if len(args) != 0 {
		return shim.Error("1111111----Incorrect number of arguments. Private marble data must be passed in transient map.")
	}

	logger.Infof("11111111111111111111111111")

	transMap, err := APIstub.GetTransient()
	if err != nil {
		return shim.Error("222222 -Error getting transient: " + err.Error())
	}

	ArtisteDataAsBytes, ok := transMap["Artiste"]
	if !ok {
		return shim.Error("Artiste must be a key in the transient map")
	}
	logger.Infof("********************8   " + string(ArtisteDataAsBytes))

	if len(ArtisteDataAsBytes) == 0 {
		return shim.Error("333333 -marble value in the transient map must be a non-empty JSON string")
	}

	logger.Infof("2222222")

	var ArtisteInput ArtisteTransientInput
	err = json.Unmarshal(ArtisteDataAsBytes, &ArtisteInput)
	if err != nil {
		return shim.Error("44444 -Failed to decode JSON of: " + string(ArtisteDataAsBytes) + "Error is : " + err.Error())
	}

	logger.Infof("3333")

	if len(ArtisteInput.Key) == 0 {
		return shim.Error("name field must be a non-empty string")
	}
	if len(ArtisteInput.Album) == 0 {
		return shim.Error("color field must be a non-empty string")
	}
	if len(ArtisteInput.Ext) == 0 {
		return shim.Error("Ext field must be a non-empty string")
	}
	if len(ArtisteInput.Title) == 0 {
		return shim.Error("color field must be a non-empty string")
	}
	if len(ArtisteInput.Owner) == 0 {
		return shim.Error("owner field must be a non-empty string")
	}
	if len(ArtisteInput.Price) == 0 {
		return shim.Error("price field must be a non-empty string")
	}

	logger.Infof("444444")

	// ==== Check if Artiste already exists ====
	ArtisteAsBytes, err := APIstub.GetPrivateData("collectionArtistes", ArtisteInput.Key)
	if err != nil {
		return shim.Error("Failed to get marble: " + err.Error())
	} else if ArtisteAsBytes != nil {
		fmt.Println("This Artiste already exists: " + ArtisteInput.Key)
		return shim.Error("This Artiste already exists: " + ArtisteInput.Key)
	}

	logger.Infof("55555")

	var Artiste = Album{AlbumName: ArtisteInput.Album, Ext: ArtisteInput.Ext, Title: ArtisteInput.Title, Owner: ArtisteInput.Owner}

	ArtisteAsBytes, err = json.Marshal(Artiste)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = APIstub.PutPrivateData("collectionArtistes", ArtisteInput.Key, ArtisteAsBytes)
	if err != nil {
		logger.Infof("6666666")
		return shim.Error(err.Error())
	}

	ArtistePrivateDetails := &artistePrivateDetails{Owner: ArtisteInput.Owner, Price: ArtisteInput.Price}

	ArtistePrivateDetailsAsBytes, err := json.Marshal(ArtistePrivateDetails)
	if err != nil {
		logger.Infof("77777")
		return shim.Error(err.Error())
	}

	err = APIstub.PutPrivateData("collectionArtistePrivateDetails", ArtisteInput.Key, ArtistePrivateDetailsAsBytes)
	if err != nil {
		logger.Infof("888888")
		return shim.Error(err.Error())
	}

	return shim.Success(ArtisteAsBytes)
}

func (s *SmartContract) updatePrivateAlbum(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	type ArtisteTransientInput struct {
		Owner string `json:"owner"`
		Title string `json:"title"`
		Key   string `json:"key"`
	}
	if len(args) != 0 {
		return shim.Error("1111111----Incorrect number of arguments. Private marble data must be passed in transient map.")
	}

	logger.Infof("11111111111111111111111111")

	transMap, err := APIstub.GetTransient()
	if err != nil {
		return shim.Error("222222 -Error getting transient: " + err.Error())
	}

	ArtisteDataAsBytes, ok := transMap["Artiste"]
	if !ok {
		return shim.Error("Artiste must be a key in the transient map")
	}
	logger.Infof("********************8   " + string(ArtisteDataAsBytes))

	if len(ArtisteDataAsBytes) == 0 {
		return shim.Error("333333 -marble value in the transient map must be a non-empty JSON string")
	}

	logger.Infof("2222222")

	var ArtisteInput ArtisteTransientInput
	err = json.Unmarshal(ArtisteDataAsBytes, &ArtisteInput)
	if err != nil {
		return shim.Error("44444 -Failed to decode JSON of: " + string(ArtisteDataAsBytes) + "Error is : " + err.Error())
	}

	ArtistePrivateDetails := &artistePrivateDetails{Owner: ArtisteInput.Owner, Price: ArtisteInput.Title}

	ArtistePrivateDetailsAsBytes, err := json.Marshal(ArtistePrivateDetails)
	if err != nil {
		logger.Infof("77777")
		return shim.Error(err.Error())
	}

	err = APIstub.PutPrivateData("collectionArtistePrivateDetails", ArtisteInput.Key, ArtistePrivateDetailsAsBytes)
	if err != nil {
		logger.Infof("888888")
		return shim.Error(err.Error())
	}

	return shim.Success(ArtistePrivateDetailsAsBytes)

}

func (s *SmartContract) createAlbum(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var Artiste = Album{AlbumName: args[1], Ext: args[2], Title: args[3], Owner: args[4]}

	ArtisteAsBytes, _ := json.Marshal(Artiste)
	APIstub.PutState(args[0], ArtisteAsBytes)

	indexName := "owner~key"
	colorNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{Artiste.Owner, args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	APIstub.PutState(colorNameIndexKey, value)

	return shim.Success(ArtisteAsBytes)
}

func (S *SmartContract) queryAlbumsByOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}
	owner := args[0]

	ownerAndIdResultIterator, err := APIstub.GetStateByPartialCompositeKey("owner~key", []string{owner})
	if err != nil {
		return shim.Error(err.Error())
	}

	defer ownerAndIdResultIterator.Close()

	var i int
	var id string

	var Artistes []byte
	bArrayMemberAlreadyWritten := false

	Artistes = append([]byte("["))

	for i = 0; ownerAndIdResultIterator.HasNext(); i++ {
		responseRange, err := ownerAndIdResultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		objectType, compositeKeyParts, err := APIstub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}

		id = compositeKeyParts[1]
		assetAsBytes, err := APIstub.GetState(id)

		if bArrayMemberAlreadyWritten == true {
			newBytes := append([]byte(","), assetAsBytes...)
			Artistes = append(Artistes, newBytes...)

		} else {
			// newBytes := append([]byte(","), ArtistesAsBytes...)
			Artistes = append(Artistes, assetAsBytes...)
		}

		fmt.Printf("Found a asset for index : %s asset id : ", objectType, compositeKeyParts[0], compositeKeyParts[1])
		bArrayMemberAlreadyWritten = true

	}

	Artistes = append(Artistes, []byte("]")...)

	return shim.Success(Artistes)
}

func (s *SmartContract) queryAllAlbums(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "Artiste0"
	endKey := "Artiste999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllArtistes:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) restictedMethod(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// get an ID for the client which is guaranteed to be unique within the MSP
	//id, err := cid.GetID(APIstub) -

	// get the MSP ID of the client's identity
	//mspid, err := cid.GetMSPID(APIstub) -

	// get the value of the attribute
	//val, ok, err := cid.GetAttributeValue(APIstub, "attr1") -

	// get the X509 certifite of the client, or nil if the client's identity was not based on an X509 certifiPratite
	//cert, err := cid.GetX509Certificate(APIstub) -

	val, ok, err := cid.GetAttributeValue(APIstub, "role")
	if err != nil {
		// There was an error trying to retrieve the attribute
		shim.Error("Error while retriving attributes")
	}
	if !ok {
		// The client identity does not possess the attribute
		shim.Error("Client identity doesnot posses the attribute")
	}
	// Do something with the value of 'val'
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return shim.Error("Only user with role as APPROVER have access this method!")
	} else {
		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}

		ArtisteAsBytes, _ := APIstub.GetState(args[0])
		return shim.Success(ArtisteAsBytes)
	}

}

func (s *SmartContract) changeAlbumOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	ArtisteAsBytes, _ := APIstub.GetState(args[0])
	Artiste := Album{}

	json.Unmarshal(ArtisteAsBytes, &Artiste)
	Artiste.Owner = args[1]

	ArtisteAsBytes, _ = json.Marshal(Artiste)
	APIstub.PutState(args[0], ArtisteAsBytes)

	return shim.Success(ArtisteAsBytes)
}

func (t *SmartContract) getHistoryForAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ArtisteName := args[0]

	resultsIterator, err := stub.GetHistoryForKey(ArtisteName)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForAsset returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) createPrivateAlbumImplicitForOrg1(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect arguments. Expecting 5 arguments")
	}

	var Artiste = Album{AlbumName: args[1], Ext: args[2], Title: args[3], Owner: args[4]}

	ArtisteAsBytes, _ := json.Marshal(Artiste)
	// APIstub.PutState(args[0], ArtisteAsBytes)

	err := APIstub.PutPrivateData("_implicit_org_Org1MSP", args[0], ArtisteAsBytes)
	if err != nil {
		return shim.Error("Failed to add asset: " + args[0])
	}
	return shim.Success(ArtisteAsBytes)
}

func (s *SmartContract) createPrivateAlbumImplicitForOrg2(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect arguments. Expecting 5 arguments")
	}

	var Artiste = Album{AlbumName: args[1], Ext: args[2], Title: args[3], Owner: args[4]}

	ArtisteAsBytes, _ := json.Marshal(Artiste)
	APIstub.PutState(args[0], ArtisteAsBytes)

	err := APIstub.PutPrivateData("_implicit_org_Org2MSP", args[0], ArtisteAsBytes)
	if err != nil {
		return shim.Error("Failed to add asset: " + args[0])
	}
	return shim.Success(ArtisteAsBytes)
}

func (s *SmartContract) queryPrivateDataHash(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	ArtisteAsBytes, _ := APIstub.GetPrivateDataHash(args[0], args[1])
	return shim.Success(ArtisteAsBytes)
}

// func (s *SmartContract) CreateArtisteAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
// 	if len(args) != 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 1")
// 	}

// 	var Artiste Artiste
// 	err := json.Unmarshal([]byte(args[0]), &Artiste)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	ArtisteAsBytes, err := json.Marshal(Artiste)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	err = APIstub.PutState(Artiste.ID, ArtisteAsBytes)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	return shim.Success(nil)
// }

// func (s *SmartContract) addBulkAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
// 	logger.Infof("Function addBulkAsset called and length of arguments is:  %d", len(args))
// 	if len(args) >= 500 {
// 		logger.Errorf("Incorrect number of arguments in function CreateAsset, expecting less than 500, but got: %b", len(args))
// 		return shim.Error("Incorrect number of arguments, expecting 2")
// 	}

// 	var eventKeyValue []string

// 	for i, s := range args {

// 		key :=s[0];
// 		var Artiste = Artiste{Album: s[1], Ext: s[2], Title: s[3], Owner: s[4]}

// 		eventKeyValue = strings.SplitN(s, "#", 3)
// 		if len(eventKeyValue) != 3 {
// 			logger.Errorf("Error occured, Please Album sure that you have provided the array of strings and each string should be  in \"EventType#Key#Value\" format")
// 			return shim.Error("Error occured, Please Album sure that you have provided the array of strings and each string should be  in \"EventType#Key#Value\" format")
// 		}

// 		assetAsBytes := []byte(eventKeyValue[2])
// 		err := APIstub.PutState(eventKeyValue[1], assetAsBytes)
// 		if err != nil {
// 			logger.Errorf("Error coocured while putting state for asset %s in APIStub, error: %s", eventKeyValue[1], err.Error())
// 			return shim.Error(err.Error())
// 		}
// 		// logger.infof("Adding value for ")
// 		fmt.Println(i, s)

// 		indexName := "Event~Id"
// 		eventAndIDIndexKey, err2 := APIstub.CreateCompositeKey(indexName, []string{eventKeyValue[0], eventKeyValue[1]})

// 		if err2 != nil {
// 			logger.Errorf("Error coocured while putting state in APIStub, error: %s", err.Error())
// 			return shim.Error(err2.Error())
// 		}

// 		value := []byte{0x00}
// 		err = APIstub.PutState(eventAndIDIndexKey, value)
// 		if err != nil {
// 			logger.Errorf("Error coocured while putting state in APIStub, error: %s", err.Error())
// 			return shim.Error(err.Error())
// 		}
// 		// logger.Infof("Created Composite key : %s", eventAndIDIndexKey)

// 	}

// 	return shim.Success(nil)
// }

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
