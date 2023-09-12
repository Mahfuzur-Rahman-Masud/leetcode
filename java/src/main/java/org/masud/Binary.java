package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class Binary {


  public int hammingWeight(int n) {
    int count =0;

    while (n != 0){
      count += n &1;
      n>>>=1;
    }

    return count;
  }


  public int reverseBitsIterative(int n){
    int result = 0,isolated_single_bit;
    for (int i = 0; i <32; i++) {
      isolated_single_bit = (n >>> i) & 1;
      result = result | isolated_single_bit << (31-i);
    }

    return result;
  }

  public int reverseBits(int n){
    // 1234_5678
    n = (n & 0xffff_0000) >>>16 | (n& 0x0000_ffff)<<16; // 1234_0000>>> | 0000_5678<< = 5678_1234
    n = (n & 0xff_00_ff_00) >>>8 | (n& 0x00_ff_00_ff)<<8; // 56_00_12_00>>> | 00_78_00_34<< = 78_56_34_12
    n = (n & 0xf0_f0_f0_f0) >>>4 | (n& 0x0f_0f_0f_0f)<<4; // 70_50_30_10>>> | 08_07_04_02<< = 87_65_43_21
    n = (n & 0xcccccccc) >>>2 | (n& 0x33333333)<<2; //  [7]0111 & 1100>>>2 | [7]0111 & 0011<<2 =[13]11_01
    n = (n & 0xaaaaaaaa) >>>1 | (n& 0x55555555)<<1; //  [13]1101 & 1010>>>2 | [13]1101 & 0101<<2 =[14]11_10
    return n;
  }

  public int reverseBits_v1(int n) {
    StringBuilder binaryString = new StringBuilder(Integer.toBinaryString(n));

      while (binaryString.length() < 32){
        binaryString.insert(0, '0');
      }

    byte[] binary = binaryString.toString().getBytes();


    int markIn = 0, markOut = binary.length - 1;
    byte current;
    while (markIn < markOut) {
      current = binary[markIn];
      binary[markIn] = binary[markOut];
      binary[markOut] = current;
      markIn++;
      markOut--;
    }

    return (int) Long.parseLong(new String(binary), 2);
  }


  public String addBinary(String a, String b) {
    StringBuilder stringBuilder = new StringBuilder();

    int carry = 48;
    int la = a.length() - 1; // hello 4
    int lb = b.length() - 1; // jimi 3
    int lm = Math.max(la, lb); // 4
    int offA = lm - la;
    int offB = lm - lb;
    int ia, ib;

    for (int i = lm; i >= 0; i--) {
      ia = i - offA;
      ib = i - offB;

      int s = (ia < 0 ? 48 : a.charAt(ia)) + (ib < 0 ? 48 : b.charAt(ib)) + carry;
      // 144 000
      // 145 001
      // 146 011
      // 147 111

      if (s == 144 || s == 146) {
        stringBuilder.append('0');
      } else if (s == 145 || s == 147) {
        stringBuilder.append('1');
      }

      if (s > 145) {
        carry = 49;
      } else {
        carry = 48;
      }
    }

    if (carry == 49) {
      stringBuilder.append('1');
    }

    return stringBuilder.reverse().toString();

  }

  public String addBinaryFailsForTooLongString(String a, String b) {
    int i = Integer.parseInt(a, 2);
    int j = Integer.parseInt(b, 2);
    return Integer.toBinaryString(i + j);
  }

}
