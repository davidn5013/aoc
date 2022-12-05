My solution

Python was better for this...

Short solution in golang same speed and about same file size... 
1404kb or 1342kb but that it with the input file embed

	Solution 1 = (528) 528
	2022/12/04 23:12:51 took 4.1761ms
	Solution 2 = (881) 881
	2022/12/04 23:12:51 took 4.7324ms
	609.3611ms run and compile
	
	Part A: 528
	Part B: 881
	took   8.435ms
	605.3248ms run and compile

	Solution 1 = (528) 528
	2022/12/04 23:16:47 took 4.1972ms
	Solution 2 = (881) 881
	2022/12/04 23:16:47 took 4.2103ms
	144.6807ms execute 
	
	Part A: 528
	Part B: 881
	took  9.7953ms
	107.0728ms execute

python 3.11
function "part1" took 2 ms
528
function "part2" took 1 ms
881
41.3036ms - 46.2702ms extra time is python interpreter 
so total is about 48ms that hart get exact golang 100-144 ms 
