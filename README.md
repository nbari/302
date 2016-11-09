# 302
🐸  redirect service


Design considerations


DB boltdb:

- If you require a high random write throughput (>10,000 w/sec) or you need to
use spinning disks then LevelDB could be a good choice. If your application is
read-heavy or does a lot of range scans then Bolt could be a good choice.
