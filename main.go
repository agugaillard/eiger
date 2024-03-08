package main

import (
	"errors"
	"os"
	"strconv"

	"encoding/json"

	"github.com/agugaillard/eiger/rdiff"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	differ *rdiff.Rdiff

	errBadInvocation = errors.New("program not invoked properly")
)

func init() {
	logger, _ = zap.NewDevelopment()
	differ = rdiff.NewDefaultRdiff(logger)
}

func main() {
	if len(os.Args) == 1 {
		panic(errBadInvocation)
	}
	function := os.Args[1]
	var err error
	switch function {
	case "signature":
		err = signature()
	case "delta":
		err = delta()
	case "patch":
		err = patch()
	default:
		err = errBadInvocation
	}
	if err != nil {
		panic(err)
	}
}

// Args:
// 0 program
// 1 function
// 2 input
// 3 output
// 4 chunksize
func signature() error {
	if len(os.Args) < 4 || len(os.Args) > 5 {
		return errBadInvocation
	}
	input, err := os.ReadFile(os.Args[2])
	if err != nil {
		return err
	}
	var chunksize uint
	if len(os.Args) == 5 {
		chunksizeParameter, err := strconv.Atoi(os.Args[4])
		if err != nil {
			return err
		}
		chunksize = uint(chunksizeParameter)
	}
	signature := differ.Signature(input, chunksize)
	signatureBytes, _ := json.Marshal(signature)
	return os.WriteFile(os.Args[3], signatureBytes, 0644)
}

// Args:
// 0 program
// 1 function
// 2 input
// 3 signature
// 4 output
func delta() error {
	if len(os.Args) != 5 {
		return errBadInvocation
	}
	input, err := os.ReadFile(os.Args[2])
	if err != nil {
		return err
	}
	signatureFile, err := os.ReadFile(os.Args[3])
	if err != nil {
		return err
	}
	var signature rdiff.Signature
	if err = json.Unmarshal(signatureFile, &signature); err != nil {
		return err
	}
	delta := differ.Delta(input, signature)
	deltaBytes, _ := json.Marshal(delta)
	return os.WriteFile(os.Args[4], deltaBytes, 0644)
}

// Args:
// 0 program
// 1 function
// 2 input
// 3 delta
// 4 output
func patch() error {
	if len(os.Args) != 5 {
		return errBadInvocation
	}
	input, err := os.ReadFile(os.Args[2])
	if err != nil {
		return err
	}
	deltaFile, err := os.ReadFile(os.Args[3])
	if err != nil {
		return err
	}
	var delta rdiff.Delta
	if err = json.Unmarshal(deltaFile, &delta); err != nil {
		return err
	}
	output := differ.Patch(input, delta)
	return os.WriteFile(os.Args[4], []byte(output), 0644)
}
