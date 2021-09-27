<?php

$f = fopen('in2.txt', 'w+');

$c = 1000000;
fwrite($f, "1\n");
fwrite($f, "$c\n");

for ($i = 0; $i < $c; $i++) {
    $tag = uniqid('');
    fwrite($f, "<$tag>\n");
}
fclose($f);