// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_file_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/qnsoft/common/os/qn_time"
	qn_conv "github.com/qnsoft/common/util/qn_conv"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_IsDir(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		paths := "/testfile"
		createDir(paths)
		defer delTestFiles(paths)

		t.Assert(qn_file.IsDir(testpath()+paths), true)
		t.Assert(qn_file.IsDir("./testfile2"), false)
		t.Assert(qn_file.IsDir("./testfile/tt.txt"), false)
		t.Assert(qn_file.IsDir(""), false)
	})
}

func Test_IsEmpty(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		path := "/testdir_" + qn_conv.String(qn_time.TimestampNano())
		createDir(path)
		defer delTestFiles(path)

		t.Assert(qn_file.IsEmpty(testpath()+path), true)
		t.Assert(qn_file.IsEmpty(testpath()+path+qn_file.Separator+"test.txt"), true)
	})
	qn_test.C(t, func(t *qn_test.T) {
		path := "/testfile_" + qn_conv.String(qn_time.TimestampNano())
		createTestFile(path, "")
		defer delTestFiles(path)

		t.Assert(qn_file.IsEmpty(testpath()+path), true)
	})
	qn_test.C(t, func(t *qn_test.T) {
		path := "/testfile_" + qn_conv.String(qn_time.TimestampNano())
		createTestFile(path, "1")
		defer delTestFiles(path)

		t.Assert(qn_file.IsEmpty(testpath()+path), false)
	})
}

func Test_Create(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			err       error
			filepaths []string
			fileobj   *os.File
		)
		filepaths = append(filepaths, "/testfile_cc1.txt")
		filepaths = append(filepaths, "/testfile_cc2.txt")
		for _, v := range filepaths {
			fileobj, err = qn_file.Create(testpath() + v)
			defer delTestFiles(v)
			fileobj.Close()
			t.Assert(err, nil)
		}
	})
}

func Test_Open(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)

		file1 := "/testfile_nc1.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)

		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "./testfile/file1/c1.txt")
		flags = append(flags, false)

		for k, v := range files {
			fileobj, err = qn_file.Open(testpath() + v)
			fileobj.Close()
			if flags[k] {
				t.Assert(err, nil)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_OpenFile(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)

		files = append(files, "./testfile/file1/nc1.txt")
		flags = append(flags, false)

		f1 := "/testfile_tt.txt"
		createTestFile(f1, "")
		defer delTestFiles(f1)

		files = append(files, f1)
		flags = append(flags, true)

		for k, v := range files {
			fileobj, err = qn_file.OpenFile(testpath()+v, os.O_RDWR, 0666)
			fileobj.Close()
			if flags[k] {
				t.Assert(err, nil)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_OpenWithFlag(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)

		file1 := "/testfile_t1.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)
		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "/testfiless/dirfiles/t1_no.txt")
		flags = append(flags, false)

		for k, v := range files {
			fileobj, err = qn_file.OpenWithFlag(testpath()+v, os.O_RDWR)
			fileobj.Close()
			if flags[k] {
				t.Assert(err, nil)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_OpenWithFlagPerm(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)
		file1 := "/testfile_nc1.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)
		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "/testfileyy/tt.txt")
		flags = append(flags, false)

		for k, v := range files {
			fileobj, err = qn_file.OpenWithFlagPerm(testpath()+v, os.O_RDWR, 666)
			fileobj.Close()
			if flags[k] {
				t.Assert(err, nil)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_Exists(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		var (
			flag  bool
			files []string
			flags []bool
		)

		file1 := "/testfile_GetContents.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)

		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "./testfile/havefile1/tt_no.txt")
		flags = append(flags, false)

		for k, v := range files {
			flag = qn_file.Exists(testpath() + v)
			if flags[k] {
				t.Assert(flag, true)
			} else {
				t.Assert(flag, false)
			}

		}

	})
}

func Test_Pwd(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		paths, err := os.Getwd()
		t.Assert(err, nil)
		t.Assert(qn_file.Pwd(), paths)

	})
}

func Test_IsFile(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			flag  bool
			files []string
			flags []bool
		)

		file1 := "/testfile_tt.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)
		files = append(files, file1)
		flags = append(flags, true)

		dir1 := "/testfiless"
		createDir(dir1)
		defer delTestFiles(dir1)
		files = append(files, dir1)
		flags = append(flags, false)

		files = append(files, "./testfiledd/tt1.txt")
		flags = append(flags, false)

		for k, v := range files {
			flag = qn_file.IsFile(testpath() + v)
			if flags[k] {
				t.Assert(flag, true)
			} else {
				t.Assert(flag, false)
			}

		}

	})
}

func Test_Info(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			err    error
			paths  string = "/testfile_t1.txt"
			files  os.FileInfo
			files2 os.FileInfo
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)
		files, err = qn_file.Info(testpath() + paths)
		t.Assert(err, nil)

		files2, err = os.Stat(testpath() + paths)
		t.Assert(err, nil)

		t.Assert(files, files2)

	})
}

func Test_Move(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths     string = "/ovetest"
			filepaths string = "/testfile_ttn1.txt"
			topath    string = "/testfile_ttn2.txt"
		)
		createDir("/ovetest")
		createTestFile(paths+filepaths, "a")

		defer delTestFiles(paths)

		yfile := testpath() + paths + filepaths
		tofile := testpath() + paths + topath

		t.Assert(qn_file.Move(yfile, tofile), nil)

		// 检查移动后的文件是否真实存在
		_, err := os.Stat(tofile)
		t.Assert(os.IsNotExist(err), false)

	})
}

func Test_Rename(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths  string = "/testfiles"
			ypath  string = "/testfilettm1.txt"
			topath string = "/testfilettm2.txt"
		)
		createDir(paths)
		createTestFile(paths+ypath, "a")
		defer delTestFiles(paths)

		ypath = testpath() + paths + ypath
		topath = testpath() + paths + topath

		t.Assert(qn_file.Rename(ypath, topath), nil)
		t.Assert(qn_file.IsFile(topath), true)

		t.AssertNE(qn_file.Rename("", ""), nil)

	})

}

func Test_DirNames(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths    string = "/testdirs"
			err      error
			readlist []string
		)
		havelist := []string{
			"t1.txt",
			"t2.txt",
		}

		// 创建测试文件
		createDir(paths)
		for _, v := range havelist {
			createTestFile(paths+"/"+v, "")
		}
		defer delTestFiles(paths)

		readlist, err = qn_file.DirNames(testpath() + paths)

		t.Assert(err, nil)
		t.AssertIN(readlist, havelist)

		_, err = qn_file.DirNames("")
		t.AssertNE(err, nil)

	})
}

func Test_Glob(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths      string = "/testfiles/*.txt"
			dirpath    string = "/testfiles"
			err        error
			resultlist []string
		)

		havelist1 := []string{
			"t1.txt",
			"t2.txt",
		}

		havelist2 := []string{
			testpath() + "/testfiles/t1.txt",
			testpath() + "/testfiles/t2.txt",
		}

		//===============================构建测试文件
		createDir(dirpath)
		for _, v := range havelist1 {
			createTestFile(dirpath+"/"+v, "")
		}
		defer delTestFiles(dirpath)

		resultlist, err = qn_file.Glob(testpath()+paths, true)
		t.Assert(err, nil)
		t.Assert(resultlist, havelist1)

		resultlist, err = qn_file.Glob(testpath()+paths, false)

		t.Assert(err, nil)
		t.Assert(formatpaths(resultlist), formatpaths(havelist2))

		_, err = qn_file.Glob("", true)
		t.Assert(err, nil)

	})
}

func Test_Remove(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths string = "/testfile_t1.txt"
		)
		createTestFile(paths, "")
		t.Assert(qn_file.Remove(testpath()+paths), nil)

		t.Assert(qn_file.Remove(""), nil)

		defer delTestFiles(paths)

	})
}

func Test_IsReadable(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1 string = "/testfile_GetContents.txt"
			paths2 string = "./testfile_GetContents_no.txt"
		)

		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		t.Assert(qn_file.IsReadable(testpath()+paths1), true)
		t.Assert(qn_file.IsReadable(paths2), false)

	})
}

func Test_IsWritable(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1 string = "/testfile_GetContents.txt"
			paths2 string = "./testfile_GetContents_no.txt"
		)

		createTestFile(paths1, "")
		defer delTestFiles(paths1)
		t.Assert(qn_file.IsWritable(testpath()+paths1), true)
		t.Assert(qn_file.IsWritable(paths2), false)

	})
}

func Test_Chmod(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1 string = "/testfile_GetContents.txt"
			paths2 string = "./testfile_GetContents_no.txt"
		)
		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		t.Assert(qn_file.Chmod(testpath()+paths1, 0777), nil)
		t.AssertNE(qn_file.Chmod(paths2, 0777), nil)

	})
}

// 获取绝对目录地址
func Test_RealPath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1    string = "/testfile_files"
			readlPath string

			tempstr string
		)

		createDir(paths1)
		defer delTestFiles(paths1)

		readlPath = qn_file.RealPath("./")

		tempstr, _ = filepath.Abs("./")

		t.Assert(readlPath, tempstr)

		t.Assert(qn_file.RealPath("./nodirs"), "")

	})
}

// 获取当前执行文件的目录
func Test_SelfPath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1    string
			readlPath string
			tempstr   string
		)
		readlPath = qn_file.SelfPath()
		readlPath = filepath.ToSlash(readlPath)

		tempstr, _ = filepath.Abs(os.Args[0])
		paths1 = filepath.ToSlash(tempstr)
		paths1 = strings.Replace(paths1, "./", "/", 1)

		t.Assert(readlPath, paths1)

	})
}

func Test_SelfDir(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1    string
			readlPath string
			tempstr   string
		)
		readlPath = qn_file.SelfDir()

		tempstr, _ = filepath.Abs(os.Args[0])
		paths1 = filepath.Dir(tempstr)

		t.Assert(readlPath, paths1)

	})
}

func Test_Basename(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1    string = "/testfilerr_GetContents.txt"
			readlPath string
		)

		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		readlPath = qn_file.Basename(testpath() + paths1)
		t.Assert(readlPath, "testfilerr_GetContents.txt")

	})
}

func Test_Dir(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1    string = "/testfiless"
			readlPath string
		)
		createDir(paths1)
		defer delTestFiles(paths1)

		readlPath = qn_file.Dir(testpath() + paths1)

		t.Assert(readlPath, testpath())

	})
}

func Test_Ext(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1   string = "/testfile_GetContents.txt"
			dirpath1        = "/testdirs"
		)
		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		createDir(dirpath1)
		defer delTestFiles(dirpath1)

		t.Assert(qn_file.Ext(testpath()+paths1), ".txt")
		t.Assert(qn_file.Ext(testpath()+dirpath1), "")
	})

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_file.Ext("/var/www/test.js"), ".js")
		t.Assert(qn_file.Ext("/var/www/test.min.js"), ".js")
		t.Assert(qn_file.Ext("/var/www/test.js?1"), ".js")
		t.Assert(qn_file.Ext("/var/www/test.min.js?v1"), ".js")
	})
}

func Test_ExtName(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_file.ExtName("/var/www/test.js"), "js")
		t.Assert(qn_file.ExtName("/var/www/test.min.js"), "js")
		t.Assert(qn_file.ExtName("/var/www/test.js?v=1"), "js")
		t.Assert(qn_file.ExtName("/var/www/test.min.js?v=1"), "js")
	})
}

func Test_TempDir(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		if qn_file.Separator != "/" || !qn_file.Exists("/tmp") {
			t.Assert(qn_file.TempDir(), os.TempDir())
		} else {
			t.Assert(qn_file.TempDir(), "/tmp")
		}
	})
}

func Test_Mkdir(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			tpath string = "/testfile/createdir"
			err   error
		)

		defer delTestFiles("/testfile")

		err = qn_file.Mkdir(testpath() + tpath)
		t.Assert(err, nil)

		err = qn_file.Mkdir("")
		t.AssertNE(err, nil)

		err = qn_file.Mkdir(testpath() + tpath + "2/t1")
		t.Assert(err, nil)

	})
}

func Test_Stat(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			tpath1   = "/testfile_t1.txt"
			tpath2   = "./testfile_t1_no.txt"
			err      error
			fileiofo os.FileInfo
		)

		createTestFile(tpath1, "a")
		defer delTestFiles(tpath1)

		fileiofo, err = qn_file.Stat(testpath() + tpath1)
		t.Assert(err, nil)

		t.Assert(fileiofo.Size(), 1)

		_, err = qn_file.Stat(tpath2)
		t.AssertNE(err, nil)

	})
}

func Test_MainPkgPath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		reads := qn_file.MainPkgPath()
		t.Assert(reads, "")
	})
}
