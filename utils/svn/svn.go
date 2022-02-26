package svn

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var svn *SVN
var localAddr string

func Init() error {
	localAddr = viper.GetString("svn.local_addr")

	var err error
	user := viper.GetString("svn.user")
	password := viper.GetString("svn.password")

	opts := &Options{Timeout: 300 * time.Second, Username: user, Password: password, NonInteractive: true, NoAuthCache: true, TrustServerCertFailures: "unknown-ca,cn-mismatch,expired,not-yet-valid,other"}

	svn, err = NewSVN(localAddr, opts)
	if err != nil {
		fmt.Println("svn NewSVN err", err)
		return err
	}
	return nil
}

func AddSvn(filepath string) error {

	if err := svn.Add(filepath, nil); err != nil {
		fmt.Println("svn add err", err, filepath)
		return err
	}

	fmt.Println("svn add ok!")
	return nil
}

func CommitSvn() {
	var err error
	if err = svn.Update(localAddr, nil); err != nil {
		fmt.Println("svn update err", err, localAddr)
		return
	}
	if err = svn.Commit(localAddr, "update file", nil); err != nil {
		fmt.Println("svn commit err", err, localAddr)
		return
	}
	fmt.Println("svn commit ok!")
}
