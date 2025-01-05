# shellcheck disable=SC2086
docker build -t bogdanserdinov/chart-playground:$1 .
docker push bogdanserdinov/chart-playground:"$1"
