package org.masud;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.function.BiFunction;
import java.util.function.Function;

/**
 * @author Mahfuzur Rahman
 */
public class Eval2<K1, K2,  V> {
  private static class Case<K1,K2,V>{
    public Case(K1 k1, K2 k2, V v) {
      this.k1 = k1;
      this.k2 = k2;
      this.v = v;
    }

    private final K1 k1;
    private final K2 k2;
    private final V v;
    private V r;
  }
  private final BiFunction<K1, K2,V> fun;

  private final List<Case<K1, K2,V>> cases = new ArrayList<>();

  public Eval2(BiFunction<K1,K2, V> fun) {
    this.fun = fun;
  }

  public Eval2<K1,K2,V> put(K1 k1, K2 k2, V v){
    cases.add(new Case<>(k1, k2, v));
    return this;
  }


  public void evaluate(){
    for (Case<K1,K2,V> c : cases) {
      K1 k1 = c.k1;
      K2 k2 = c.k2;
      V v = c.v;
      c.r = fun.apply(k1, k2);
      System.out.printf("K1: %s\tK2: %s\t V: %s\t R:%s%n", k1,k2, v, c.r);
      if(v==null && c.r==null){
        continue;
      }

      if(v!= null && c.r!=null && v.equals(c.r)){
        continue;
      }

      throw new RuntimeException("Test failed");
    }
  }
}
