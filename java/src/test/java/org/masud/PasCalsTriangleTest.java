package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import java.util.List;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class PasCalsTriangleTest extends PasCalsTriangle{

  @Test
  void generateTest() {
    List<List<Integer>> list = generate(5);
    System.out.printf("Out: %s%n", list);
  }
}