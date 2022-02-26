package svn

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
	"time"
)

const localSVNPath = "/Users/hehaitao/temp/Go"

var localSVN *SVN
var svnUrl string
var workDir string

func init() {
	//svnUrl = "https://120.25.76.199:10008/svn/CarbonChain/tags/CarCo2/GO"
	workDir = "/Users/hehaitao/temp/Go"

	//fmt.Println("=----=", svnUrl, workDir)
	var err error
	opts := &Options{Timeout: 300 * time.Second}
	localSVN, err = NewSVN(localSVNPath, opts)
	if err != nil {
		fmt.Println("----111-----", err)
		return
		//log.Fatal(err)
	}
	//InitLocalRepo()
}

func InitLocalRepo() {
	if err := localSVN.Checkout("", localSVNPath); err != nil {
		fmt.Println("----100-----")
		log.Fatal(err)
	}
	//if err := localSVN.Mkdir(DefaultBranchesDir); err != nil {
	//	fmt.Println("----101-----")
	//	log.Fatal(err)
	//}
	//if err := localSVN.Mkdir(DefaultTrunkDir); err != nil {
	//	fmt.Println("----102-----")
	//	log.Fatal(err)
	//}
	//if err := localSVN.Mkdir(DefaultTagsDir); err != nil {
	//	fmt.Println("----103-----")
	//	log.Fatal(err)
	//}
}

func TestNewSVN(t *testing.T) {
	//workDir = "/Users/hehaitao/temp/Go"
	//fmt.Println("=----=", svnUrl, workDir)
	var err error
	opts := &Options{WorkDir: workDir, Timeout: 300 * time.Second}
	localSVN, err = NewSVN(localSVNPath, opts)
	if err != nil {
		fmt.Println("----111-----", err)
		return
		//log.Fatal(err)
	}
}

func TestSVNAddCommit(t *testing.T) {

	fn := path.Join(workDir, "sample5.txt")
	fp, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0644)
	fp.WriteString("# SVN Test File Fuck YOU 22!\n")
	if err != nil {
		fmt.Println("333", err)
		return
		//t.Fatal(err)
	}
	defer fp.Close()
	if err := localSVN.Add(fn, nil); err != nil {
		fmt.Println("444", err)
		//t.Fatal(err)
		//return
	}
	if err := localSVN.Commit(workDir, "add sample.txt", nil); err != nil {
		fmt.Println("555", err)
		//t.Fatal(err)
	}
}

func TestSVNAddCommit2(t *testing.T) {
	fn := path.Join(workDir, localSVNPath, DefaultTrunkDir, "sample2.txt")
	fp, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0644)
	fp.WriteString("# SVN Test File Line\n")
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()
	if err := localSVN.Commit(localSVNPath, "update sample.txt", nil); err != nil {
		t.Fatal(err)
	}
}
func TestCleanUp(t *testing.T) {
	if err := localSVN.Cleanup(localSVNPath); err != nil {
		t.Error(err)
	}
}

func inArray(a []string, str string) bool {
	for _, s := range a {
		if str == s {
			return true
		}
	}
	return false
}

func TestNewBranch(t *testing.T) {
	bn := "develop"
	if err := svn.NewBranch(bn, "create branch develop"); err != nil {
		t.Fatal(err)
	}
	if branches, err := svn.Branches(); err != nil {
		t.Error(err)
	} else if !inArray(branches, bn) {
		t.Error("new branch fail, not exists")
	}
}
func TestNewTag(t *testing.T) {
	tn := "v0.1"
	if err := svn.NewTag(tn, "create tag v0.1"); err != nil {
		t.Fatal(err)
	}
	if tages, err := svn.Tags(); err != nil {
		t.Error(err)
	} else if !inArray(tages, tn) {
		t.Error("new tag fail, not exists")
	}
}

func TestLog(t *testing.T) {
	lr, err := svn.Log(path.Join(svn.trunkDir, "sample.txt"))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(lr)

	if len(lr.Logentrys) != 2 {
		t.Errorf("log error expect 2 logentry get %d", len(lr.Logentrys))
	}
}

func TestExport(t *testing.T) {
	err := svn.Export("trunk/sample.txt", "/tmp/sample.txt", "-r", "1")
	if err != nil {
		t.Fatal(err)
	}
	fp, err := os.Open("/tmp/sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "# SVN Test File\n" {
		t.Error("Export File Err")
	}
}
