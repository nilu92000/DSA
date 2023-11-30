import java.util.*;  
public class ne
{
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		System.out.println("Java Program to print all even numbers in array ");
        System.out.print("Enter the size of array: ");
        int size = sc.nextInt();
        int arr[] = new int[size];
        for(int i=0; i<size; i++) {
            System.out.print("Please give value for index "+ i +" : ");
            arr[i] = sc.nextInt();
        }
        System.out.println("Even number in array are");
        for(int i=0; i<size; i++)
        {
            if(arr[i]%2==0){
                System.out.print(arr[i]+"\t");
            }
        }
	}
}