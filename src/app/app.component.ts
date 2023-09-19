import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'testyaml';
  ngOnInit(): void {
    //localStorage.setItem('TryIt_securitySchemeValues', JSON.stringify({'OAuth2Demo': 'Bearer AwesomeDemo', 'OAuth2Prod': 'Bearer AwesomeProd'}))
  }
}
