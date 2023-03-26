# ceph 교육

### Cluster 부팅

```
# ceph orch ls
# ceph -s
# ceph osd unset noout
# osd unset norecover
# ceph osd unset norebalance
# ceph osd unset nobackfill
# ceph osd unset nodown
# ceph osd unset pause
```

Ceph File System(`CephFS`)을 사용하는 경우 `joinable` 플래그를 `true` 로 설정하여 `CephFS` 클러스터를 백업합니다.

```
# ceph fs set cephfs joinable true
```



### 모니터링

```
# podman exec -it ceph-mon-MONITOR_NAME /bin/bash
# podman exec -it ceph-499829b4-832f-11eb-8d6d-001a4a000635-mon.host01 /bin/bash
```



```
# cephadm shell
# ceph health
# ceph status 
# ceph -w
```




