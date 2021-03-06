
| io interface    | 嵌套 interface               | func                                                         | 功能                                                         | buffer 是否实现 |
| --------------- | ---------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | --------------- |
| Reader          |                              | Read<font color="silver">(p []byte) (n int, err error)</font> | 用于输出自身的数据（-> p）                                   | ✅               |
| Writer          |                              | Write<font color="silver">(p []byte) (n int, err error)</font> | 用于将数据存入自身（p ->）                                   | ✅               |
| Closer          |                              | Close<font color="silver">() error</font>                    | 用于关闭数据读写（关闭文件、通道、连接、数据库……）           |                 |
| Seeker          |                              | Seek<font color="silver">(offset int64, whence int) (int64, error)</font> | 用于移动数据的读写指针（每一次读写操作都从指针位置开始）     |                 |
| ReadWriter      | Reader<br>Writer             |                                                              | 组合接口                                                     | ✅               |
| ReadCloser      | Reader<br>Closer             |                                                              | 组合接口                                                     |                 |
| WriteCloser     | Writer<br/>Closer            |                                                              | 组合接口                                                     |                 |
| ReadWriteCloser | Reader<br/>Writer<br/>Closer |                                                              | 组合接口                                                     |                 |
| ReadSeeker      | Reader<br/>Seeker            |                                                              | 组合接口                                                     |                 |
| ReadSeekCloser  | Reader<br/>Seeker<br/>Closer |                                                              | 组合接口                                                     |                 |
| WriteSeeker     | Writer<br/>Seeker            |                                                              | 组合接口                                                     |                 |
| ReadWriteSeeker | Reader<br/>Writer<br/>Seeker |                                                              | 组合接口                                                     |                 |
| ReaderFrom      |                              | ReadFrom<font color="silver">(r Reader) (n int64, err error)</font> | 用于从 r 中读取数据存入自身                                  | ✅               |
| WriterTo        |                              | WriteTo<font color="silver">(w Writer) (n int64, err error)</font> | 用于将自身的数据写入 w 中                                    | ✅               |
| ReaderAt        |                              | ReadAt<font color="silver">(p []byte, off int64) (n int, err error)</font> | 用于从指定偏移位置开始，输出自身的数据（-> p）               |                 |
| WriterAt        |                              | WriteAt<font color="silver">(p []byte, off int64) (n int, err error)</font> | 用于从指定偏移位置开始，将数据存入自身（p ->）               |                 |
| ByteReader      |                              | ReadByte<font color="silver">() (byte, error)</font>         | 用于从自身读出一个字节                                       | ✅               |
| ByteScanner     | ByteReader                   | UnreadByte<font color="silver">() error</font>               | 用于从自身读出一个字节，且可以撤销最后一次读取（下次可以读出一样的数据） | ✅               |
| ByteWriter      |                              | WriteByte<font color="silver">(c byte) error</font>          | 用于将一个字节写入自身                                       | ✅               |
| RuneReader      |                              | ReadRune<font color="silver">() (r rune, size int, err error)</font> | 用于从自身读取一个 UTF-8 编码的字符到 r 中                   | ✅               |
| RuneScanner     | RuneReader                   | UnreadRune<font color="silver">() error</font>               | 用于从自身读取一个 UTF-8 编码的字符到 r 中，且可以撤销最后一次读取（下次可以读出一样的数据） | ✅               |
| StringWriter    |                              | WriteString<font color="silver">(s string) (n int, err error)</font> | 用于将字符串 s 写入到 w 中                                   | ✅               |



| ioutil 方法                                                  | 功能                                     | Go 1.16 之后推荐替代方法 |
| ------------------------------------------------------------ | ---------------------------------------- | ------------------------ |
| ReadAll<font color="silver">(r io.Reader) ([]byte, error)</font> | 从 io.Reader 中读取全部数据，返回 []byte | io.ReadAll               |
| ReadDir<font color="silver">(dirname string) ([]fs.FileInfo, error)</font> | 读取目录下的所有文件（向下一级）         | os.ReadDir               |
| ReadFile<font color="silver">(filename string) ([]byte, error)</font> | 读取文件，返回 []byte                    | os.ReadFile              |
| WriteFile<font color="silver">(filename string, data []byte, perm fs.FileMode) error</font> | 把 []byte 写到文件                       | os.WriteFile             |
| NopCloser<font color="silver">(r io.Reader) io.ReadCloser</font> | 把 io.Reader 包装成 io.ReadCloser        | io.NopCloser             |
| TempDir<font color="silver">(dir, pattern string) (name string, err error)</font> | 创建临时目录                             | --                       |
| TempFile<font color="silver">(dir, pattern string) (f *os.File, err error)</font> | 创建临时文件                             | --                       |

