package main.java.com.pjs.shoppingcart;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class ShoppingCart {
    public static void main(String[] args) {
        Map<String, Integer> priceList = new HashMap<>();
        priceList.put("Apple", 35);
        priceList.put("Banana", 20);
        priceList.put("Melon", 50);
        priceList.put("Lime", 15);

        List<String> shoppingBasket = new ArrayList<>();
        shoppingBasket.add("Apple");
        shoppingBasket.add("Apple");
        shoppingBasket.add("Banana");
        shoppingBasket.add("Melon");
        shoppingBasket.add("Melon");
        shoppingBasket.add("Lime");
        shoppingBasket.add("Lime");
        shoppingBasket.add("Lime");

        int totalCost = calculateTotalCost(shoppingBasket, priceList);
        System.out.println("Total cost of the shopping basket: " + totalCost + "p");
    }

    public static int calculateTotalCost(List<String> shoppingBasket, Map<String, Integer> priceList) {
        Map<String, Integer> itemCount = new HashMap<>();
        int totalCost = 0;

        for (String item : shoppingBasket) {
            itemCount.put(item, itemCount.getOrDefault(item, 0) + 1);
        }

        for (String item : itemCount.keySet()) {
            int count = itemCount.get(item);
            int price = priceList.get(item);

            if (item.equals("Melon")) {
                // Buy one get one free for Melons
                totalCost += (count / 2 + count % 2) * price;
            } else if (item.equals("Lime")) {
                // Three for the price of two for Limes
                totalCost += ((count / 3) * 2 + count % 3) * price;
            } else {
                totalCost += count * price;
            }
        }

        return totalCost;
    }
}
