<?php
include "bucket.php";

function benchmark($x)
{
    $total = 1e6;
    $now = microtime(true);
    for ($i = 0; $i < $total; $i++) {
        $x();
    }
    $end = microtime(true);
    return ($end - $now) / $total;
}

$bucket = MyBucket::getMe();

$benchmark = benchmark(function () use ($bucket) {
    $bucket->get("123");
});

echo number_format($benchmark * 1e9, 1) . "ns \n";