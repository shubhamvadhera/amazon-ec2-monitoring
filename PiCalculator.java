public class PiCalculator {
    static double i; // Control variable
    static double s = 1; // Signal for the next iteration
    static double pi = 0;

    public static void main(String[] args) {
        System.out.println(
                "Approximation of the number PI through the Leibniz's series\n");
        for (i = 1;; i += 2) {
            pi = pi + s * (4 / i);
            s = -s;
            System.out.println(pi);
        }
    }
}
