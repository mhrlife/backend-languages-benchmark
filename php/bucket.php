<?php


class MyBucket
{
    private $db;
    private static $me;

    static public function getMe()
    {
        if (self::$me == null) {

            $db = [];
            for ($i = 0; $i < 1e6; $i++) {
                array_push($db, ["key" => $i, "value" => "value"]);
            }
            self::$me = new MyBucket($db);
        }
        return self::$me;
    }

    public function __construct($db)
    {
        $this->db = $db;
    }

    public function get(string $key)
    {
        $len = count($this->db);
        if ($len == 0) {
            return false;
        }
        $low = 0;
        $high = $len - 1;
        while ($low <= $high) {
            $mid = floor(($high + $low) / 2);
            $item = $this->db[$mid];
            $val = $item["key"];
            if ($val == $key)
                return $item;
            if ($key < $val)
                $high = $mid - 1;
            if ($key > $val)
                $low = $mid + 1;
        }
        return null;
    }
}

