# 常用排序算法，PHP版本

## 冒泡排序
> 冒泡排序：循环n-1组，相邻的元素从左到右，两两比较
```
  //冒泡排序
  function bubbleSort($arr)
  {
      $length = count($arr);
      for ($i = 0; $i < $length - 1; $i++) {
          for ($j = 0; $j < $length - 1 - $i; $j++) {
              if ($arr[$j] > $arr[$j + 1]) {
                  $tmp = $arr[$j];
                  $arr[$j] = $arr[$j + 1];
                  $arr[$j + 1] = $tmp;
              }
          }
      }
      print_r($arr);
  }
```

## 选择排序
> 选择排序：循环n-1组，从左到右，每个元素与其之后的元素比较，不断比较并更新当前值为最小值
```
  //选择排序
  function selectSort($arr)
  {
      $length = count($arr);
      for ($i = 0; $i < $length - 1; $i++) {
          for ($j = $i + 1; $j < $length; $j++) {
              if ($arr[$i] > $arr[$j]) {
                  $tmp = $arr[$i];
                  $arr[$i] = $arr[$j];
                  $arr[$j] = $tmp;
              }
          }
      }
      print_r($arr);
  }
```

## 插入排序
> 插入排序：循环从无序的数组中取值，插入到有序的列表中，假定第一个元素是有序数组，所以循环从第二个元素开始
``` 
  //插入排序
  function insertSort($arr)
  {
      $length = sizeof($arr);
      for ($i = 1; $i < $length; $i++) {
          $preIndex = $i - 1;
          $current = $arr[$i];
          while ($preIndex >= 0 && $current < $arr[$preIndex]) {
              $arr[$preIndex + 1] = $arr[$preIndex];
              $preIndex--;
          }
          $arr[$preIndex + 1] = $current;
      }
      print_r($arr);
  }
```

## 归并排序
> 归并排序：分治思想，按照数组位置不断分割直到不可再分，将最后的数据合并
```
  //归并排序主程序
  function msort(&$arr, $low, $high)
  {
      if ($low >= $high) {
          return;
      }
      $mid = floor(($low + $high) / 2);
      msort($arr, $low, $mid);
      msort($arr, $mid + 1, $high);
      mmerge($arr, $low, $mid, $high);
  }

  //合并操作
  function mmerge(&$arr, $low, $mid, $high)
  {
      //复制数据备用
      $arr1 = [];
      for ($k = $low; $k <= $high; $k++) {
          $arr1[$k] = $arr[$k];
      }

      //合并数组
      $i = $low;
      $j = $mid + 1;
      for ($k = $low; $k <= $high; $k++) {
          if ($i > $mid) {
              $arr[$k] = $arr1[$j++];
          } elseif ($j > $high) {
              $arr[$k] = $arr1[$i++];
          } elseif ($arr1[$i] > $arr1[$j]) {
              $arr[$k] = $arr1[$j++];
          } else {
              $arr[$k] = $arr1[$i++];
          }
      }
  }
```

## 快速排序
> 与归并排序不同的是，归并依赖中间位置的数据，而快速排序则依赖数字大小进行划分，默认取第一个数字作为划分标准，最后合并
```
  //快速排序
  function quickSort($arr)
  {
      if (count($arr) < 1) {
          return $arr;
      }
      $spliteData = $arr[0];
      $left = $right = $mid = [];
      foreach ($arr as $key => $value) {
          if ($value > $spliteData) {
              $right[] = $value;
          } elseif ($value < $spliteData) {
              $left[] = $value;
          } else {
              $mid[] = $value;
          }
      }

      $left = quickSort($left);
      $right = quickSort($right);
      $newArr = array_merge($left, $mid, $right);
      return $newArr;
  }
```
