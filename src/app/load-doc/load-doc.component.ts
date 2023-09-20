import { AfterViewInit, Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-load-doc',
  templateUrl: './load-doc.component.html',
  styleUrls: ['./load-doc.component.scss']
})
export class LoadDocComponent implements OnInit {
  public showContent: boolean = false;

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {

    setTimeout(()=>this.showContent=true, 1000);
    
    this.apiService.generateYaml().subscribe((response:any) => {

      console.log("DONE with generation!!")
      
    });
  }

  ngOnDestroy(): void {
  }

}
