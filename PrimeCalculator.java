class PrimeCalculator {
    public static void main(String args[]) {
        boolean status;
        long num = 2;
        while (true) {
            status = true;
            for (long j = 2; j < num; j++) {
                if (num % j == 0) {
                    status = false;
                    break;
                }
            }
            if (status) {
                System.out.println(num);
            }
            num++;
        }
    }
}