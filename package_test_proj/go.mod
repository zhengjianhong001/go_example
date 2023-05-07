module package_test_proj

go 1.16

require package_test_proj2/mypkg2 v0.0.0

replace package_test_proj2/mypkg2 => ../package_test_proj2
