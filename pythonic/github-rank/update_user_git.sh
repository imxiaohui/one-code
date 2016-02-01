user=$1
dist_dir=$2
CWD=$(cd `dirname $0` ; pwd)
gits=`curl "https://api.github.com/users/${user}/repos" 2> /dev/null | grep "full_name" | awk -F ":|\""  '{print $5}' | awk -F "/" '{print $2}'`

cd $dist_dir
for g in $gits;do 
    if [ -d "$g" ];then
        cd $g 
        git pull
        cd - 
    else
        git clone git@github.com:${user}/${g}.git                
    fi
done
cd $CWD
