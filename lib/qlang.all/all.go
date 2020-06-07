package qall

import (
	"github.com/xushiwei/qlang/lib/bufio"
	"github.com/xushiwei/qlang/lib/bytes"
	"github.com/xushiwei/qlang/lib/crypto/md5"
	"github.com/xushiwei/qlang/lib/encoding/hex"
	"github.com/xushiwei/qlang/lib/encoding/json"
	"github.com/xushiwei/qlang/lib/eqlang"
	"github.com/xushiwei/qlang/lib/errors"
	"github.com/xushiwei/qlang/lib/io"
	"github.com/xushiwei/qlang/lib/io/ioutil"
	"github.com/xushiwei/qlang/lib/math"
	"github.com/xushiwei/qlang/lib/meta"
	"github.com/xushiwei/qlang/lib/net/http"
	"github.com/xushiwei/qlang/lib/os"
	"github.com/xushiwei/qlang/lib/path"
	"github.com/xushiwei/qlang/lib/reflect"
	"github.com/xushiwei/qlang/lib/runtime"
	"github.com/xushiwei/qlang/lib/strconv"
	"github.com/xushiwei/qlang/lib/strings"
	"github.com/xushiwei/qlang/lib/sync"
	"github.com/xushiwei/qlang/lib/terminal"
	"github.com/xushiwei/qlang/lib/tpl/extractor"
	"github.com/xushiwei/qlang/lib/version"
	qlang "github.com/xushiwei/qlang/spec"

	// qlang builtin modules
	_ "github.com/xushiwei/qlang/lib/builtin"
	_ "github.com/xushiwei/qlang/lib/chan"
)

// -----------------------------------------------------------------------------

// Copyright prints qlang copyright information.
//
func Copyright() {
	version.Copyright()
}

// InitSafe inits qlang and imports modules.
//
func InitSafe(safeMode bool) {

	qlang.SafeMode = safeMode

	qlang.Import("", math.Exports) // import math as builtin package
	qlang.Import("", meta.Exports) // import meta package
	qlang.Import("bufio", bufio.Exports)
	qlang.Import("bytes", bytes.Exports)
	qlang.Import("md5", md5.Exports)
	qlang.Import("io", io.Exports)
	qlang.Import("ioutil", ioutil.Exports)
	qlang.Import("hex", hex.Exports)
	qlang.Import("json", json.Exports)
	qlang.Import("errors", errors.Exports)
	qlang.Import("eqlang", eqlang.Exports)
	qlang.Import("math", math.Exports)
	qlang.Import("os", os.Exports)
	qlang.Import("", os.InlineExports)
	qlang.Import("path", path.Exports)
	qlang.Import("http", http.Exports)
	qlang.Import("reflect", reflect.Exports)
	qlang.Import("runtime", runtime.Exports)
	qlang.Import("strconv", strconv.Exports)
	qlang.Import("strings", strings.Exports)
	qlang.Import("sync", sync.Exports)
	qlang.Import("terminal", terminal.Exports)
	qlang.Import("extractor", extractor.Exports)
}

// -----------------------------------------------------------------------------
