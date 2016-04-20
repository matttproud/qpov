// -*- html -*-
package main

const doneTmpl = `
{{ $root := . }}
<h2>Finished</h2>
<table>
  <tr>
    <th>Order</th>
    <th>Lease</th>
    <th>Created</th>
    <th>Done</th>
    <th>Time</th>
    <th>Image</th>
    <!--  <th>Package</th> -->
    <th>File</th>
    <!--  <th>Args</th> -->
    <th>Address</th>
    <th>Hostname</th>
  </tr>
  {{range .DoneLeases}}
  <tr>
    <td nowrap class="fixed">{{.OrderId}}</td>
    <td nowrap class="fixed"><a href="{{$root.Root}}/lease/{{.LeaseId}}">{{.LeaseId}}</a></td>
    <td nowrap>{{.CreatedMs|fmsdate "2006-01-02 15:04"}}</td>
    <td nowrap>{{.UpdatedMs|fmsdate "2006-01-02 15:04"}}</td>
    <td nowrap class="fixed">{{.CreatedMs|fmssub .UpdatedMs}}</td>
    <td nowrap><a href="{{$root.Root}}/image/{{.LeaseId}}">Image</a></td>
    <!--  <td nowrap>{{.Order.Package|fileonly}}</td> -->
    <td nowrap>{{.Order.File}}</td>
    <td nowrap>{{.Address}}</td>
    <td nowrap>{{.Hostname}}</td>
    <!--  <td nowrap>{{.Order.Args}}</td> -->
  </tr>
  {{end}}
</table>
`
