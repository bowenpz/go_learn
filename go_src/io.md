





| io interface    | 嵌套 interface               | func                                            | 功能                                                         | buffer 是否实现 |
| --------------- | ---------------------------- | ----------------------------------------------- | ------------------------------------------------------------ | --------------- |
| Reader          |                              | Read(p []byte) (n int, err error)               | 用于输出自身的数据（-> p）                                   | ✅               |
| Writer          |                              | Write(p []byte) (n int, err error)              | 用于将数据存入自身（p ->）                                   | ✅               |
| Closer          |                              | Close() error                                   | 用于关闭数据读写（关闭文件、通道、连接、数据库……）           |                 |
| Seeker          |                              | Seek(offset int64, whence int) (int64, error)   | 用于移动数据的读写指针（每一次读写操作都从指针位置开始）     |                 |
| ReadWriter      | Reader<br>Writer             |                                                 | 组合接口                                                     | ✅               |
| ReadCloser      | Reader<br>Closer             |                                                 | 组合接口                                                     |                 |
| WriteCloser     | Writer<br/>Closer            |                                                 | 组合接口                                                     |                 |
| ReadWriteCloser | Reader<br/>Writer<br/>Closer |                                                 | 组合接口                                                     |                 |
| ReadSeeker      | Reader<br/>Seeker            |                                                 | 组合接口                                                     |                 |
| ReadSeekCloser  | Reader<br/>Seeker<br/>Closer |                                                 | 组合接口                                                     |                 |
| WriteSeeker     | Writer<br/>Seeker            |                                                 | 组合接口                                                     |                 |
| ReadWriteSeeker | Reader<br/>Writer<br/>Seeker |                                                 | 组合接口                                                     |                 |
| ReaderFrom      |                              | ReadFrom(r Reader) (n int64, err error)         | 用于从 r 中读取数据存入自身                                  | ✅               |
| WriterTo        |                              | WriteTo(w Writer) (n int64, err error)          | 用于将自身的数据写入 w 中                                    | ✅               |
| ReaderAt        |                              | ReadAt(p []byte, off int64) (n int, err error)  | 用于从指定偏移位置开始，输出自身的数据（-> p）               |                 |
| WriterAt        |                              | WriteAt(p []byte, off int64) (n int, err error) | 用于从指定偏移位置开始，将数据存入自身（p ->）               |                 |
| ByteReader      |                              | ReadByte() (byte, error)                        | 用于从自身读出一个字节                                       | ✅               |
| ByteScanner     | ByteReader                   | UnreadByte() error                              | 用于从自身读出一个字节，且可以撤销最后一次读取（下次可以读出一样的数据） | ✅               |
| ByteWriter      |                              | WriteByte(c byte) error                         | 用于将一个字节写入自身                                       | ✅               |
| RuneReader      |                              | ReadRune() (r rune, size int, err error)        | 用于从自身读取一个 UTF-8 编码的字符到 r 中                   | ✅               |
| RuneScanner     | RuneReader                   | UnreadRune() error                              | 用于从自身读取一个 UTF-8 编码的字符到 r 中，且可以撤销最后一次读取（下次可以读出一样的数据） | ✅               |
| StringWriter    |                              | WriteString(s string) (n int, err error)        | 用于将字符串 s 写入到 w 中                                   | ✅               |





