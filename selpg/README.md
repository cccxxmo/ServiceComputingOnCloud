# selpg

## 设计说明
整个selpg程序分为五个部分：

 - `selpg_args{...}`: 声明包含命令参数的结构体
 - `usage()`: 在出现错误时打印帮助信息
 - `process_args()`: 将输入的命令处理为`selpg_args`结构体中的参数，并在输入不正确时报告错误
 - `process_input()`: 设置输入、输出，执行相关的文件操作和打印操作
 - `main()`: 主函数。初始化参数，调用`process_args()`和`process_input()`

## 使用
见[开发 Linux 命令行实用程序 ](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)

## 测试结果

 1.  `$ selpg -s1 -e1 input_file`  
![](https://i.loli.net/2019/09/18/fF4GC2akgndl98U.png)
 2.  `$ selpg -s1 -e1 < input_file`  
![](https://i.loli.net/2019/09/18/eLW9kHoUVZpbirE.png)
 3.  `$ other_command | selpg -s1 -e2`  
![](https://i.loli.net/2019/09/18/4bEmMjuUrSLsXzO.png)
 4.  `$ selpg -s2 -e3 input_file >output_file`  
![](https://i.loli.net/2019/09/18/287iLUzVm6Hwo54.png)  
![](https://i.loli.net/2019/09/18/CDrSzQlRGyO5M4o.png)
 5.  `$ selpg -s0 -e1 input_file 2>error_file`  
![](https://i.loli.net/2019/09/18/57I6aGFDKluXSHw.png)  
![](https://i.loli.net/2019/09/18/MLTBU6HIvb1KSpg.png)
 6.  `$ selpg -s2 -e1 input_file >output_file 2>error_file`  
![](https://i.loli.net/2019/09/18/F96IajWdrRtACcu.png)  
![](https://i.loli.net/2019/09/18/lZ92CFt4wdsBKuk.png)  
![](https://i.loli.net/2019/09/18/8nSe5XCIyWMtubp.png)
 7.  `$ selpg -s2 -e1 input_file >output_file 2>/dev/null`  
![](https://i.loli.net/2019/09/18/BvcEzlgsj8Aqfrx.png)  
![](https://i.loli.net/2019/09/18/HsKDbmLJ7w6rWSU.png)
 8.  `$ selpg -s1 -e2 input_file >/dev/null`  
`$ selpg -s-1 -e2 input_file >/dev/null`  
![](https://i.loli.net/2019/09/18/xQgjBleyCRJF6hr.png)  
![](https://i.loli.net/2019/09/18/Sgv3KiUr5e7IJqC.png)  
 9.  `$ selpg -s1 -e2 input_file | other_command`  
![](https://i.loli.net/2019/09/18/QjDyCXpJMPsoU2m.png)  
 10.  `$selpg -s1 -e0 input_file 2>error_file | other_command`  
![](https://i.loli.net/2019/09/18/zwWkfOcAZDhq3Ja.png)  
![](https://i.loli.net/2019/09/18/almh82LAu7ojfY5.png)  
 11.  `$ selpg -s3 -e4 -l3 input_file`  
![](https://i.loli.net/2019/09/18/wNfk5W1IapnAYm2.png)  
 12.  `$ selpg -s2 -e2 -f input_file`  
![](https://i.loli.net/2019/09/18/Pc7kR13BgsFODEa.png)  
 13.  `$ selpg -s10 -e20 -dlp1 input_file`  
因虚拟机上无打印机设备，所以该命令暂时无法测试。
