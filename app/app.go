package app

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	ipfs "github.com/ipfs/go-ipfs-api"
)

type Uploader interface {
	Upload(ctx context.Context, filepath string) error
	Store(ctx context.Context, cid string) error
}

type uploader struct {
	ipfsShell *ipfs.Shell
	ethClient *ethclient.Client
}

func New(ipfsShell ipfs.Shell, ethClient ethclient.Client) *uploader {
	return &uploader{ipfsShell: ipfsShell, ethClient: ethClient}
}

func (u *uploader) Upload(ctx context.Context, filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return error
	}
	//sh := ipfs.NewShell("localhost:5001")
	cid, err := u.ipfsShell.Add(file)
	if err != nil {
		return error
	}
	return cid, nil
}

func (u *uploader) Store(ctx contetx.Context, cid string) error {

}
