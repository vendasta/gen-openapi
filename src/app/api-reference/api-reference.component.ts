import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-api-reference',
  templateUrl: './api-reference.component.html',
  styleUrls: ['./api-reference.component.scss']
})
export class ApiReferenceComponent implements OnInit {

  constructor(private router: Router) {}

  ngOnInit(): void {
  }

  onGenerate($event: any) {
  }


}
