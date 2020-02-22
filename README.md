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
