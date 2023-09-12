

class SmallestInfiniteSet {
  int _next = 1;

  var _list = [];

  int popSmallest() {
    if (_list.isEmpty || _next < _list.last) return _next++;
    if (_next == _list.last) _next++;
    return _list.removeLast();
  }

  void addBack(int num) {
    if (num >= _next) return;
    var set = _list.toSet()..add(num);
    _list = set.toList()..sort((a, b) => b - a);
  }
}
