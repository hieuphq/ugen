#### ugen generate uuid and copy into clipboard

```
go get github.com/hieuphq/ugen
cd $GOPATH/src/github.com/hieuphq/ugen
make install
```

#### Usage command

- Generate an uuid string

```
ugen
```

- Save 300 uuid(s) into clipboard

```
ugen -a 300
```

#### Checklists

- [x] Generate uuid
- [x] Generate uuid with amount
- [x] Copy into clipboard
- [x] Support MacOS
- [ ] Support Linux