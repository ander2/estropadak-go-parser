# Estropadak Go Parser

Simple estropada parser (Euskolabel, Euskotrenm, ARC1 and ARC2) built to learn and test Go lang.

## Current usage

It parses an HTML result and prints the final classification as text table. By default, it read the content from Stdin, but an URL can be provided with the `-u` flag.

The `-t` flag specifies wich kind of parser (ACT or ARC) to use.

```
$ ./estropadakparser -h
Usage of estropadakparser:
  -t string
        Parser type: ACT or ARC (default "ACT")
  -u string
        Content URL: http://www.liga-arc.com/es/regata/489/xvii.-hondarribiko-arrantzaleen-kofradiko-bandera
```

### Parsing stdin
```
$ cat html/zarautz_act_2023_1.html | go run .
   XLVI. Zarauzko Ikurriña (J1) (19-08-2023)
1  BERMEO URDAIBAI                 [04:55 09:57 15:16] 20:12,14
2  ORIO ORIALKI                    [04:55 09:59 15:16] 20:20,30
3  AMENABAR DONOSTIARRA UR KIROLAK [04:53 10:03 15:27] 20:30,40
4  HONDARRIBIA                     [04:56 10:05 15:32] 20:32,92
5  ZIERBENA BAHIAS DE BIZKAIA      [04:58 10:08 15:31] 20:38,46
6  GETARIA                         [05:01 10:14 15:35] 20:49,76
7  CR CABO DA CRUZ                 [05:02 10:21 15:46] 21:07,90
8  LEKITTARRA ELECNOR              [05:06 10:24 15:53] 21:11,08
9  ITSASOKO AMA SANTURTZI          [05:12 10:29 15:58] 21:16,36
10 SAMERTOLAMEU FANDICOSTA         [05:11 10:32 16:05] 21:22,44
11 ONDARROA CIKAUTXO               [05:11 10:32 16:05] 21:29,16
12 KAIKU BEREZ GALANTA             [05:13 10:39 16:11] 21:31,20
```

### Fetching content and parsing
```
$./estropadakparser.go -u "https://www.euskolabelliga.com/resultados/ver.php?id=es&r=1678275014" -t ACT
   XLVI. Zarauzko Ikurriña (J1) (19-08-2023)
1  BERMEO URDAIBAI                 [04:55 09:57 15:16] 20:12,14
2  ORIO ORIALKI                    [04:55 09:59 15:16] 20:20,30
3  AMENABAR DONOSTIARRA UR KIROLAK [04:53 10:03 15:27] 20:30,40
4  HONDARRIBIA                     [04:56 10:05 15:32] 20:32,92
5  ZIERBENA BAHIAS DE BIZKAIA      [04:58 10:08 15:31] 20:38,46
6  GETARIA                         [05:01 10:14 15:35] 20:49,76
7  CR CABO DA CRUZ                 [05:02 10:21 15:46] 21:07,90
8  LEKITTARRA ELECNOR              [05:06 10:24 15:53] 21:11,08
9  ITSASOKO AMA SANTURTZI          [05:12 10:29 15:58] 21:16,36
10 SAMERTOLAMEU FANDICOSTA         [05:11 10:32 16:05] 21:22,44
11 ONDARROA CIKAUTXO               [05:11 10:32 16:05] 21:29,16
12 KAIKU BEREZ GALANTA             [05:13 10:39 16:11] 21:31,20
```

Euskotren estropadak are also supported:

```
$ cat html/fabrika_euskotren_2023.html | ./estropadakparser
  VII Bandera Fabrika (08-07-2023)
1 NORTINDAL DONOSTIARRA UR KIROLAK [05:37  ] 10:44,90
2 DONOSTIA ARRAUN LAGUNAK          [05:35  ] 10:48,98
3 ORIO ORIALKI                     [05:39  ] 10:54,10
4 TOLOSALDEA ARRAUN KLUBA          [05:46  ] 11:12,46
5 HONDARRIBIA BERTAKO IGOGAILUAK   [05:48  ] 11:12,86
6 HIBAIKA JAMONES ANCIN            [05:53  ] 11:25,96
7 CR CABO DA CRUZ                  [05:55  ] 11:30,80
8 SD TIRÁN PEREIRA                 [06:03  ] 11:50,24
```

ARC1 estropadas example:

```
$ cat html/hondarribia_arc1_2023.html | ./estroapadakparser -t ARC
   XVII. HONDARRIBIKO ARRANTZALEEN KOFRADIKO BANDERA
1  Lapurdi Antton Bilbao        [5:16 10:18 15:58] 20:47,50
2  Arkote A.  T.                [5:17 10:20 16:06] 20:59,16
3  Sanpedrotarra A.e.           [5:19 10:16 15:59] 21:00,32
4  Zarautz Gesalaga Okelan      [5:21 10:20 16:08] 21:10,00
5  Camargo                      [5:25 10:31 16:20] 21:15,56
6  Pedreña                      [5:22 10:26 16:13] 21:18,59
7  San Juan CMO Valves          [5:24 10:30 16:17] 21:19,65
8  Zumaiako Telmo Deun A.k.e..  [5:24 10:29 16:15] 21:22,77
9  Hondarribia                  [5:24 10:34 16:28] 21:37,17
10 Busturialdea                 [5:29 10:39 16:35] 21:44,97
11 Deusto-bilbao                [5:32 10:52 16:51] 22:05,25
12 Castro Canteras de Santullan [5:35 10:53 16:53] 22:09,71
```
