
import "archive/tar"
 tar包实现了tar格式压缩文件的存取。本包目标是覆盖大多数tar的变种，包括GNU和BSD生成的tar文件。
import "archive/zip"
 zip包提供了zip档案文件的读写服务。本包不支持跨硬盘的压缩。

import "bufio"
 bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。
import "builtin"
 builtin包为Go的预声明标识符提供了文档。此处列出的条目其实并不在builtin 包中，对它们的描述只是为了让 godoc 给该语言的特殊标识符提供文档。

import "bytes"
 bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似。

import "compress/bzip2"  bzip2包实现bzip2的解压缩。
import "compress/flate"
import "compress/gzip"
import "compress/lzw"
import "compress/zlib"

import "container/heap" heap包提供了对任意类型（实现了heap.Interface接口）的堆操作。（最小）堆是具有“每个节点都是以其为根的子树中最小值”属性的树。树的最小元素为其根元素，索引0的位置。heap是常用的实现优先队列的方法。要创建一个优先队列，实现一个具有使用（负的）优先级作为比较的依据的Less方法的Heap接口，如此一来可用Push添加项目而用Pop取出队列最高优先级的项目。
import "container/list"
import "container/ring"

import "context"

import "crypto" crypto包搜集了常用的密码（算法）常量。
import "crypto/aes"
import "crypto/cipher"
import "crypto/des"
import "crypto/dsa"
import "crypto/ecdsa"
import "crypto/elliptic"
import "crypto/hmac"
import "crypto/md5"
import "crypto/rand"
import "crypto/rc4"
import "crypto/rsa"
import "crypto/sha1"
import "crypto/sha256"
import "crypto/sha512"
import "crypto/subtle"
import "crypto/tls"
import "crypto/x509"
import "crypto/x509/pkix"

import "database/sql"
import "database/sql/driver"

import "debug/dwarf"
import "debug/elf"
import "debug/gosym"
import "debug/macho"
import "debug/pe"
import "debug/plan9obj"

import "encoding"  encoding包定义了供其它包使用的可以将数据在字节水平和文本表示之间转换的接口。encoding/gob、encoding/json、encoding/xml三个包都会检查使用这些接口。因此，只要实现了这些接口一次，就可以在多个包里使用。标准包内建类型time.Time和net.IP都实现了这些接口。接口是成对的，分别产生和还原编码后的数据。
import "encoding/ascii85"
import "encoding/asn1"
import "encoding/base32"
import "encoding/base64"
import "encoding/binary"
import "encoding/csv"
import "encoding/gob"
import "encoding/hex"
import "encoding/json"
import "encoding/pem"
import "encoding/xml"

import "errors"
 errors包实现了创建错误值的函数。

import "expvar"
 expvar包提供了公共变量的标准接口，如服务的操作计数器。本包通过HTTP在/debug/vars位置以JSON格式导出了这些变量。
 
import "flag"
 flag包实现了命令行参数的解析。

import "fmt"
 mt包实现了类似C语言printf和scanf的格式化I/O。格式化动作（'verb'）源自C语言但更简单。
	func Print(a ...interface{}) (n int, err error)
	func Println(a ...interface{}) (n int, err error)
	func Printf(format string, a ...interface{}) (n int, err error)
	func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	func Sprint(a ...interface{}) string
	func Sprintln(a ...interface{}) string
	func Sprintf(format string, a ...interface{}) string
	func Errorf(format string, a ...interface{}) error

	func Scan(a ...interface{}) (n int, err error)
	func Scanln(a ...interface{}) (n int, err error)
	func Scanf(format string, a ...interface{}) (n int, err error)
	func Fscan(r io.Reader, a ...interface{}) (n int, err error)
	func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
	func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	func Sscan(str string, a ...interface{}) (n int, err error)
	func Sscanln(str string, a ...interface{}) (n int, err error)
	func Sscanf(str string, format string, a ...interface{}) (n int, err error)

import "go/ast"
import "go/build"
import "go/constant"
import "go/doc"
 doc包从Go的AST提取源码文档。
import "go/format"
import "go/importer"
import "go/parser"
import "go/printer"
import "go/scanner"
import "go/token"
import "go/types"

import "hash"
 hash包提供hash函数的接口。
import "hash/adler32"
import "hash/crc32"
import "hash/crc64"
import "hash/fnv"

import "html"
 html包提供了用于转义和解转义HTML文本的函数。
import "html/template"

import "image"
import "image/color"
import "image/color/palette"
import "image/draw"
import "image/gif"
import "image/jpeg"
import "image/png"

import "index/suffixarray"
 suffixarrayb包通过使用内存中的后缀树实现了对数级时间消耗的子字符串搜索。

 
import "io"
 io包提供了对I/O原语的基本接口。本包的基本任务是包装这些原语已有的实现（如os包里的原语），使之成为共享的公共接口，这些公共接口抽象出了泛用的函数并附加了一些相关的原语的操作。因为这些接口和原语是对底层实现完全不同的低水平操作的包装，除非得到其它方面的通知，客户端不应假设它们是并发执行安全的。
import "io/ioutil"

import "log"
 log包实现了简单的日志服务。
import "log/syslog"

import "math"
 math包提供了基本的数学常数和数学函数。
import "math/big"
import "math/cmplx"
import "math/rand"

import "mime"
 mime实现了MIME的部分规定。
import "mime/multipart"
import "mime/quotedprintable"

import "net"
 net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket。
import "net/http"
import "net/http/cgi"
import "net/http/cookiejar"
import "net/http/fcgi"
import "net/http/httptest"
import "net/http/httptrace"
import "net/http/httputil"
import "net/http/pprof"
import "net/mail"
import "net/rpc"
import "net/rpc/jsonrpc"
import "net/smtp"
import "net/textproto"
import "net/url"

import "os"
 os包提供了操作系统函数的不依赖平台的接口。
import "os/exec"
import "os/signal"
import "os/user"


import "path"
 path实现了对斜杠分隔的路径的实用操作函数。
import "path/filepath"

import "plugin"

import "reflect"


import "regexp"
 regexp包实现了正则表达式搜索。
import "regexp/syntax"

import "runtime"
 runtime包提供和go运行时环境的互操作，如控制go程的函数。它也包括用于reflect包的低层次类型信息；参见reflect报的文档获取运行时类型系统的可编程接口。
import "runtime/cgo"
import "runtime/debug"
import "runtime/pprof"
import "runtime/race"
import "runtime/trace"

import "sort"
 sort包提供了排序切片和用户自定义数据集的函数。


import "strconv"
 strconv包实现了基本数据类型和其字符串表示的相互转换。

import "strings"
 strings包实现了用于操作字符的简单函数。

import "sync"
 sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。
import "sync/atomic

import "syscall"

import "testing"
 testing 提供对 Go 包的自动化测试的支持。
import "testing/iotest"
import "testing/quick"

import "text/scanner"
import "text/tabwriter"
import "text/template"
import "text/template/parse"

import "time"
 time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。

import "unicode"
import "unicode/utf16"
import "unicode/utf8"

import "unsafe"
 unsafe包提供了一些跳过go语言类型安全限制的操作。





















































