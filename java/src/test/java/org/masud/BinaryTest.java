package org.masud;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class BinaryTest extends Binary {


  @Test
  void toBinaryStr() {
    System.out.println(Integer.toBinaryString(0x44));
    assert 0x4 %4 ==0;
    assert 0x44%4==0;
    assert 0x444444%4 ==0;
    long c1 = System.currentTimeMillis();
    for (int i = 0; i < 5000000; i++) {
      assert getX(i)%4 ==0;
    }
    long c2 = System.currentTimeMillis();
    for (int i = 0; i < 5000000; i++) {
      int n = getX(i);
      assert (0xfffffffc   | n) == 0xfffffffc  ;
    }
    long c3 = System.currentTimeMillis();
    System.out.printf("A: %s\nB:%s\n", c2-c1, c3-c2);


  }

  int getX(int n){
    return n*4+4;
  }

  @Test
  void bitReverseBench() {
    long check0 = System.currentTimeMillis();

    for (int i = 0; i < 50000000; i++) {
      reverseBits(i*5321+11);
    }
    long check1 = System.currentTimeMillis();

    for (int i = 0; i < 50000000; i++) {
      reverseBitsIterative(i*5321+11);
    }
    long check2 = System.currentTimeMillis();

    System.out.printf("Normal: %s\nIterative: %s\n", check1 - check0, check2-check1);

  }

  @Test
  void bitReverseTest() {

    reverseBits(15);

    new Eval<>(this::reverseBitsIterative)
      .put(43261596, 964176192)
      .put(-3, -1073741825)


//      .put((int) Long.parseLong("11111111111111111111111111111101", 2), (int) Long.parseLong("10111111111111111111111111111111"))
      .evaluate();
      ;

    System.out.printf("%s %s%n", 1, reverseBits_v1(1));
    System.out.printf("%s %s | %s%n", 43261596, reverseBits_v1(43261596), 964176192);
    System.out.printf("%s %s | %s%n", -3, reverseBits_v1(-3), -1073741825);
  }

  @Test
  void testAddBinary() {
    new Eval2<>(this::addBinary)
      .put("1", "1", "10")
      .put("1", "0", "1")
      .put("0", "0", "0")
      .put("11", "1", "100")
      .put("1010", "1011", "10101")
//      .put("0000", "1", "1")
      .put("11", "10", "101")
      .put("11", "11", "110")
//      .put("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011", "110")
      .evaluate();
  }

  @Test
  void testAscii() {
    System.out.printf("value of %s%n", ((int) '0' ));
    System.out.printf("value of %s%n", 48*3); // 145 146 147
    // carry 146 | 147
    System.out.printf("value of %s%n", 5|3);
  }
}