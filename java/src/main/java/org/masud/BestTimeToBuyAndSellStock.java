package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class BestTimeToBuyAndSellStock {

  public int maxProfit(int[] prices) {
    if (prices == null || prices.length == 0) return 0;

    int bottom = Integer.MAX_VALUE;
    int profit = 0;
    int price = 0;
    int profitToday =0;

    for (int i = 0; i < prices.length ; i++) {
       price = prices[i];
       if(price < bottom){
         bottom = price;
       }else{
         profitToday = price -bottom;
          if(profitToday > profit){
            profit = profitToday;
          }
       }
    }

    return profit;
  }
}
