# 压测结果

### 方法调用与普通加法运算的性能对比

```text
go test -gcflags=all="-N -l" -run=. -bench=Basic -count=5
goos: darwin
goarch: amd64
pkg: github.com/xnslong/benchmark-go
BenchmarkBasic/plus-8           558065978                2.08 ns/op
BenchmarkBasic/plus-8           572090719                2.12 ns/op
BenchmarkBasic/plus-8           565462380                2.19 ns/op
BenchmarkBasic/plus-8           533984989                2.17 ns/op
BenchmarkBasic/plus-8           555679914                2.23 ns/op
BenchmarkBasic/call_func-8      468455794                2.54 ns/op
BenchmarkBasic/call_func-8      472724515                2.72 ns/op
BenchmarkBasic/call_func-8      503349939                2.81 ns/op
BenchmarkBasic/call_func-8      494187294                2.48 ns/op
BenchmarkBasic/call_func-8      497432827                2.35 ns/op
PASS
ok      github.com/xnslong/benchmark-go 14.955s
```

### 栈调用与循环迭代的性能对比

```text
go test -gcflags=all="-N -l" -run=. -bench=Stack -count=5        
goos: darwin
goarch: amd64
pkg: github.com/xnslong/benchmark-go
BenchmarkStack/A:loop-8                  2930137               411 ns/op
BenchmarkStack/A:loop-8                  2985661               401 ns/op
BenchmarkStack/A:loop-8                  2989432               402 ns/op
BenchmarkStack/A:loop-8                  3005659               393 ns/op
BenchmarkStack/A:loop-8                  3016893               397 ns/op
BenchmarkStack/B:stack-8                 2944824               412 ns/op
BenchmarkStack/B:stack-8                 2930708               409 ns/op
BenchmarkStack/B:stack-8                 2925718               407 ns/op
BenchmarkStack/B:stack-8                 2910595               409 ns/op
BenchmarkStack/B:stack-8                 2958751               409 ns/op
BenchmarkStack/C:stack2-8                2196970               559 ns/op
BenchmarkStack/C:stack2-8                2159953               556 ns/op
BenchmarkStack/C:stack2-8                2156610               552 ns/op
BenchmarkStack/C:stack2-8                2184256               553 ns/op
BenchmarkStack/C:stack2-8                2155498               559 ns/op
PASS
ok      github.com/xnslong/benchmark-go 25.492s
```

解释说明

```plantuml
package loop as "循环调用" {
    node lroot as " "
    node node1 as "++"
    node node2 as "++"
    node node3 as "++"

    lroot --> node1: call
    lroot --> node2: call
    lroot --> node3: call
}
note top of loop
总消耗=n*加法消耗 + n*调用消耗
end note

package stack as "栈调用" {
    node snode1 as "++"
    node snode2 as "++"
    node snode3 as "++"

    snode1 --> snode2: call
    snode2 --> snode3: call
}
note top of stack
总消耗=n*加法消耗 + (n-1)*调用消耗
end note


package doubleStack as "双倍消耗" {
    node dbranch1 as " "
    node dbranch2 as " "
    node dbranch3 as " "
    node dnode1 as "++"
    node dnode2 as "++"
    node dnode3 as "++"

    dbranch1 -right-> dnode1: call
    dbranch2 -right-> dnode2: call
    dbranch3 -right-> dnode3: call
    dbranch1 --> dbranch2: call
    dbranch2 --> dbranch3: call
}
note top of doubleStack
总消耗=n*加法消耗 + (2n-1)*调用消耗
end note
```