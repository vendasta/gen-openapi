import { Component } from '@angular/core';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-load-json',
  templateUrl: './load-json.component.html',
  styleUrls: ['./load-json.component.scss']
})
export class LoadJsonComponent {
  fields = '';

  constructor(private apiService: ApiService) {}
  
  ngOnInit(): void {
    this.apiService.getListFields().subscribe((response) => {
      this.fields = JSON.stringify(response, null, 2);
    });
  }

  saveFields() {
    console.log(this.fields)
    this.apiService.saveFields(this.fields).subscribe((response: any) => {
      console.log("SAVED")
    });
  }

  ngOnDestroy(): void {
  }
  
}
