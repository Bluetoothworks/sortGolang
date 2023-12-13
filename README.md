## How to use Golang sorter

In the git bash terminal use the following sample cURL to check its working

Golang sorter works in 2 ways, one sorts arrays concurrently and another in series.
Finally, it returns the sorted arrays along with the time taken to sort all of them.

Series Processing
```Gitbash
curl -X POST -H "Content-Type: application/json" -d '{"to_sort":[[3,2,1],[6,5,4],[9,8,7]]}' https://golangsorter.onrender.com/process-single
```

Concurrent Processing
```Gitbash
curl -X POST -H "Content-Type: application/json" -d '{"to_sort":[[3,2,1],[6,5,4],[9,8,7]]}' https://golangsorter.onrender.com/process-concurrent
```

We can change the value of what to sort manually as per our requirements.
