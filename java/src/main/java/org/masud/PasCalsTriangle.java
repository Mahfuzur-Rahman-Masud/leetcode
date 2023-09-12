package org.masud;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Mahfuzur Rahman
 */
public class PasCalsTriangle {

  public List<Integer> getRow(int rowIndex) {
     return generate(rowIndex + 1).get(rowIndex);
  }


  public List<List<Integer>> generate(int numRows) {
    List<List<Integer>> out = new ArrayList<>(numRows);
    if(numRows == 0) return out;

    List<Integer> previous = null;
    for (int i = 1; i < numRows +1; i++) {
      ArrayList<Integer> row = new ArrayList<>(i);
      out.add(row);
      row.add(1);

      if (i == 1){
        continue;
      }

      if(previous != null){
        for (int j = 1; j < previous.size() ; j++) {
          row.add(previous.get(j) + previous.get(j-1));
        }
      }

      row.add(1);
      previous = row;
    }

    return out;
  }
}
