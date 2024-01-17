# Estropadak Go Parser

Simple Euskolabel liga estropada parser built to learn and test Go lang.

## Current usage

Just parses from Stdin an HTML containing a result file's content and returns the final
classification.

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