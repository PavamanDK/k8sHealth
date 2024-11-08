package main
import "fmt"
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import "context"
import "k8s.io/client-go/kubernetes"
import "k8s.io/client-go/tools/clientcmd"
func main() {
  fmt.Println("Welcme to the FSS TS Main Menu")
  for{
    fmt.Println("\n1. Get Nodes")
    fmt.Println("2. Get Pods")
    fmt.Println("3. Failing Pods")
    fmt.Println("4. Memory Utilization")
    fmt.Println("5. Collect Logs")
    fmt.Println("6. Exit\n")
    fmt.Print("Please select an option from above Menu: ")
    var choice int
    fmt.Scan(&choice)
    config, _ := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
    clientset, _ := kubernetes.NewForConfig(config)
    switch choice{
      case 1:
	fmt.Println("\nPlease find the Node status\n")
        nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
        for _, node := range nodes.Items {
          fmt.Println(node.Name)
          for _, condition := range node.Status.Conditions{
            fmt.Printf("\t%s: %s\n", condition.Type, condition.Status)
          }
        } 
        break
      case 2:
        fmt.Println("Collecting Pod Details")
        pods, _ := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
        for _, pod := range pods.Items {
          fmt.Println(pod.Name)
        }
        break
      case 3:
        fmt.Println("Getting Failed Nodes")
        break
      case 4:
        fmt.Println("Getting Memory Utilization of the Nodes")
        break
      case 5:
        fmt.Println("Collecting logs")
        break
      default:
        return
    }
  }
}
