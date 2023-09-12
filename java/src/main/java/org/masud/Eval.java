package org.masud;

import java.util.HashMap;
import java.util.Map;
import java.util.Map.Entry;
import java.util.function.BiPredicate;
import java.util.function.Function;

/**
 * @author Mahfuzur Rahman
 */
public class Eval<K, V> {
  private final Map<K,V> map = new HashMap<>();
  private final Function<K,V> fun;

  public Eval(Function<K, V> fun) {
    this.fun = fun;
  }

  public Eval<K,V> put(K k, V v){
    map.put(k, v);
    return this;
  }

  private Function<K, String> kToString;
  private Function<V, String> vToString;
  public  Eval<K, V> kToString(Function<K, String> kToString){
    this.kToString = kToString;
    return this;
  }

  public  Eval<K, V> vToString(Function<V, String> vToString){
    this.vToString = vToString;
    return this;
  }

  private String kToString(K k){
    return kToString == null? k.toString() : kToString.apply(k);
  }

  private String vToString(V v){
    return vToString == null? v.toString() : vToString.apply(v);
  }

  private BiPredicate<V,V> isEqual;
  public Eval<K,V> toEqual(BiPredicate<V,V> test){
    isEqual = test;
    return  this;
  }

  private boolean isEqual(V a, V b){
    if(a == null && b == null){
      return true;
    }

    if(a == null || b == null){
      return false;
    }

    return isEqual == null? a.equals(b) : isEqual.test(a, b);
  }

  public void evaluate(){
    for (Entry<K, V> e : map.entrySet()) {
      K k = e.getKey();
      V v = e.getValue();
      V r = fun.apply(k);
      System.out.printf("\n--------------\nK: %s\nV: %s\nR: %s\n--------------\n", kToString(k), vToString(v), vToString(r));
      if(v==null && r==null){
        continue;
      }

      if(v != null && isEqual(v,r)){
        continue;
      }

      throw new RuntimeException("Test failed");
    }
  }
}
