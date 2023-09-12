import 'dart:io';
import 'dart:math';

import 'package:dart/dart.dart';
import 'package:test/test.dart';




void main() {
  test('calculate', () {
    expect(calculate(), 42);
  });

  var wa2 = work_async2(0);
  work_async(1);
  work_async(2);
  work_async(3);
  sleep(Duration(milliseconds: 100));
}

 Future<void>  work_async (int id)async{
  return Future(() {
    // sleep(Duration(microseconds: Random().nextInt(10)));
    print('Hello world! $id');
  });
}


Future<void>  work_async2 (int id)async {
  return Future(() {
    sleep(Duration(microseconds: Random().nextInt(10)));
    print('Hello world! $id');
  });
}