# genpw

Password Generator

## Usage

```
Usage of genpw:
  -c int
    	count of output (default -1)
  -n int
    	number of characters (default 16)
  -nc int
    	minimum count of numbers (default: any) (default -1)
  -sc int
    	minimum count of symbols (default: any) (default -1)
```

```
$ genpw
BqU{#SQ'Yl'rV0?! Un[$@B[R,7Br?[O0 mj:&09k'U^aV0j%` K_QBEPE;xYu.(*1X
T"rVWovu[:3`Lbs? HWVMy?Ge)#v2.O;B 2W!Y,V^!|M;h&9(i 'AKKQGS'9)jNX}J]
FJ[*)G}q$?NFtoxV )7wT?34*'OfVf1{= :.NR/!n2X1o_Vmlm PlM"@!8ab[/AB"l6
{pa<\10.#>s3gHG5 F85EPoIq]1~%iw4L ?H;oUu}I:;='1qU~ i[DU_^n?q~yo3Pi0
y,al,,w(//;`mwwX J<Y5v!i@f;b*VF8k LM'6tW>W0Kj[:_^3 $-68FJ%o:ti&Dr?K
4l>nW|iKcT(iNdDD 6l93\`&'p:?w"8cf B.;hUdDi@3MYJvs@ ]WL5[mu;-qS<$>4!
!kl)gmL)~~f>)C_6 #%X1V^WOdfm/BPY& ^RAhYJQitR0eTri5 JJT|21uI":V'q>2=
e}|$VeVkTF.=rjVV t-mS.g#2`2?&|w0X :{vuG.EWV\'"p~X? }y-<m1^=Lu3%.&O,
!,Bq6jciSh4,16oo >{*F>'l7*.;Mm-|, 3xlEBI,~BW"}J./N ,nwMb<\5-Wl}B@YF
2>$?HE&lc!$#BtVX QYPFLTv/tWBtR(3m ^%~[nVXk@ikm9|4s Od<R1t=ax.1%HApB
;-fA^26C^oWOvfh& ~*#G<1WY=7IU:b%{ o<P>x(t'xae4L-W2 sDj.!c852GGG-WA$
`F*.E,AD:v'|wW<9 |:(Y~6dFhUL':K<q )!J).*WvSlXAB)L> JNY>MGw)$D#0>BM'
w*i~L}@Q1Inw9:N[ _n/&;'!f$S1E8Y0O bJ~G1X"E"s8b@OhX ![v8"&RFVLqlv'gr
BeP)Y_,a^bfKS49R ek9m]K3gAf*24(:P 8"pLnDCg6pyoYu|q GWi9%U*nX[lKV)~1
.)4>{_uf]<~=K&j( [76FD7d^},|PI]\2 ]rtO)8&}E9!M.d>T 9:4nc3w0xFCJ^aT:
YJ[/;X5~Xo#0&k{; X?$E?w|>'\k4%=@= uKT&\y*0ThM^)#x$ m?;U!T|^[6)L,]eg

$ genpw -c 8 -sc 0
1eJXeEV74xlTe69m MdnnoBVOvS5LHA0N YRQ7oWTbOiS6dwPc 70yURuk9wviFnxkH
wR4OG2Gop8uWY05j O9lcMih6EB1mQtRp SyhLo20BKLOcWeNC HFgubFL6QyxBRnFO

$ genpw -c 8 -nc 0
``$^<Hu@:(penl*u nTg'uLQVXTh#@'S[ N_U`;rh`(y?d%aAk TCk/AfYA}^{.-!r>
\#uA@g'Cs~]$G!cY usFVQ}'Byi#na<EU n`!@RVeEUF&;}%Bm =kg|S&iTrVu~)x](

$ genpw -c 8 -nc 0 -sc 0
rdOvXawcxYAKhLwG PCkoVPVpYGSdyqrX YrfhGghXIVacNMaJ HrAOuvSEiIuhFwKv
FbNnMQJxyahiaBMn LAelrsxysoPNbtUB OnOFSjplyVCNygao XJGhELKhQTUxDbtM
```

## Installation

```
go get github.com/mattn/genpw
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
