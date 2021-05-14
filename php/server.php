<?php
include "vendor/autoload.php";
include "bucket.php";

$loop = React\EventLoop\Factory::create();

$bucket = MyBucket::getMe();
$server = new React\Http\Server($loop, function (Psr\Http\Message\ServerRequestInterface $request) use (&$bucket) {

    switch ($request->getUri()->getPath()) {
        case "/get":
        {
            $resp = "";
            for ($i = 0; $i < 50; $i++) {
                $number = rand(0, 1000);
                $value = $bucket->get($number);
                if ($value) {
                    $resp .= $value["value"];
                }
            }

            return new React\Http\Message\Response(
                200,
                array(
                    'Content-Type' => 'text/plain'
                ),
                $resp
            );

        }
        default:
            {
                return new React\Http\Message\Response(
                    404,
                    array(
                        'Content-Type' => 'text/plain'
                    ),
                    "Not Founded"
                );
            }
            break;
    }

});

$socket = new React\Socket\Server(8080, $loop);
$server->listen($socket);

echo "Server running at http://127.0.0.1:8080\n";

$loop->run();